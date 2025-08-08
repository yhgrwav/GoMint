package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/yhgrwav/GoMint/rest"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	c := rest.New(rest.WithTimeout(8*time.Second), rest.WithRetry(2, 500*time.Millisecond))

	var out map[string]any
	code, err := c.GetJSON(ctx, "https://jsonplaceholder.typicode.com/posts/1", &out, nil)
	if err != nil {
		log.Fatalf("GET error: %v (code=%d)", err, code)
	}
	fmt.Println("status:", code)
	fmt.Println("title:", out["title"])
}
