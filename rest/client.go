package rest

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"
)

type Client struct {
	http  *http.Client
	retry int           // кол-во повторов при временных ошибках
	sleep time.Duration // пауза между повторами
}

type Option func(*Client)

func WithTimeout(d time.Duration) Option {
	return func(c *Client) { c.http.Timeout = d }
}
func WithRetry(n int, sleep time.Duration) Option {
	return func(c *Client) { c.retry = n; c.sleep = sleep }
}

func New(opts ...Option) *Client {
	c := &Client{
		http:  &http.Client{Timeout: 10 * time.Second},
		retry: 0,
		sleep: 0,
	}
	for _, o := range opts {
		o(c)
	}
	return c
}

// DoJSON отправляет запрос с опциональным JSON-телом и парсит JSON-ответ в out.
func (c *Client) DoJSON(ctx context.Context, method, url string, in any, out any, headers map[string]string) (int, error) {
	var body io.Reader
	if in != nil {
		b, err := json.Marshal(in)
		if err != nil {
			return 0, err
		}
		body = bytes.NewReader(b)
	}

	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return 0, err
	}
	if in != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	attempts := c.retry + 1
	var resp *http.Response
	for i := 0; i < attempts; i++ {
		resp, err = c.http.Do(req)
		if err == nil && (resp.StatusCode < 500 && resp.StatusCode != 429) {
			break
		}
		// 5xx/429 — считаем временными, пробуем повтор
		if i < attempts-1 {
			if resp != nil {
				_ = drainAndClose(resp.Body)
			}
			select {
			case <-ctx.Done():
				return 0, ctx.Err()
			case <-time.After(c.sleep):
			}
			continue
		}
		break
	}
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if out == nil {
		// пользователь сам читает тело через DoRaw, либо ему не нужен ответ
		_, _ = io.Copy(io.Discard, resp.Body)
		return resp.StatusCode, nil
	}

	dec := json.NewDecoder(resp.Body)
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return resp.StatusCode, dec.Decode(out)
	}
	// попытаемся вытащить тело ошибки для отладки
	var e any
	_ = dec.Decode(&e)
	return resp.StatusCode, errors.New(http.StatusText(resp.StatusCode))
}

func drainAndClose(rc io.ReadCloser) error {
	_, _ = io.Copy(io.Discard, rc)
	return rc.Close()
}

// Утилиты-обёртки

func (c *Client) GetJSON(ctx context.Context, url string, out any, headers map[string]string) (int, error) {
	return c.DoJSON(ctx, http.MethodGet, url, nil, out, headers)
}

func (c *Client) PostJSON(ctx context.Context, url string, in any, out any, headers map[string]string) (int, error) {
	return c.DoJSON(ctx, http.MethodPost, url, in, out, headers)
}
