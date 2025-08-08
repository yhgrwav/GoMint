# GoMint / REST

## Задача
Простые GET/POST с JSON, таймаутами и базовыми ретраями (429/5xx).

## Быстрый старт
```go
c := rest.New(rest.WithTimeout(8*time.Second), rest.WithRetry(2, 500*time.Millisecond))
var out map[string]any
code, err := c.GetJSON(ctx, "https://jsonplaceholder.typicode.com/posts/1", &out, nil)
