package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/yhgrwav/GoMint/postgres"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	db, err := postgres.Open(ctx, postgres.Config{
		Host:     "127.0.0.1",
		Port:     15432,
		User:     "gomint",
		Password: "gomint",
		DBName:   "gomint",
		SSLMode:  "disable",
	})
	if err != nil {
		log.Fatal("connect error:", err)
	}
	defer db.Close()

	id, err := postgres.InsertUser(ctx, db, "Linus", "Torvalds", "linus@example.com")
	if err != nil {
		log.Fatal("insert error:", err)
	}
	fmt.Println("inserted id:", id)

	users, err := postgres.GetUsers(ctx, db, 10)
	if err != nil {
		log.Fatal("select error:", err)
	}
	fmt.Println("users:", users)

	if err := postgres.TxExample(ctx, db); err != nil {
		log.Fatal("tx error:", err)
	}
	fmt.Println("tx committed")
}
