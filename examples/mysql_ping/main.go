package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/yhgrwav/GoMint/mysql"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	db, err := mysql.Open(ctx, mysql.Config{
		Host:     "127.0.0.1",
		Port:     13306,
		User:     "gomint",
		Password: "gomint",
		DBName:   "gomint",
	})
	if err != nil {
		log.Fatal("connect error:", err)
	}
	defer db.Close()

	id, err := mysql.InsertUser(ctx, db, "Linus", "Torvalds", "linus@example.com")
	if err != nil {
		log.Fatal("insert error:", err)
	}
	fmt.Println("inserted id:", id)

	users, err := mysql.GetUsers(ctx, db, 10)
	if err != nil {
		log.Fatal("select error:", err)
	}
	fmt.Println("users:", users)

	if err := mysql.TxExample(ctx, db); err != nil {
		log.Fatal("tx error:", err)
	}
	fmt.Println("tx committed")
}
