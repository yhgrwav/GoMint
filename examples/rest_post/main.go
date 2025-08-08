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

	c := rest.New(rest.WithTimeout(8*time.Second), rest.WithRetry(2, 300*time.Millisecond))

	type In struct {
		Title  string `json:"title"`
		Body   string `json:"body"`
		UserID int    `json:"userId"`
	}
	var resp map[string]any
	code, err := c.PostJSON(ctx, "https://jsonplaceholder.typicode.com/posts",
		In{Title: "GoMint demo", Body: "Hello", UserID: 1},
		&resp, nil,
	)
	if err != nil {
		log.Fatalf("POST error: %v (code=%d)", err, code)
	}
	fmt.Println("status:", code)
	fmt.Println("id:", resp["id"])
}
