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
		Port:     15432, // как мы настроили в docker
		User:     "gomint",
		Password: "gomint",
		DBName:   "gomint",
		SSLMode:  "disable",
	})
	if err != nil {
		log.Fatal("connect error:", err)
	}
	defer db.Close()

	// Массовая вставка через prepared statement
	err = postgres.PreparedInsertMany(ctx, db, []postgres.User{
		{FirstName: "Ken", LastName: "Thompson", Email: "ken@example.com"},
		{FirstName: "Brian", LastName: "Kernighan", Email: "brian@example.com"},
		{FirstName: "Dennis", LastName: "Ritchie", Email: "dennis@example.com"},
	})
	if err != nil {
		log.Fatal("prepared insert error:", err)
	}
	fmt.Println("bulk insert done")

	// WHERE: достанем конкретного пользователя
	u, err := postgres.GetUserByEmail(ctx, db, "brian@example.com")
	if err != nil {
		log.Fatal("get by email error:", err)
	}
	fmt.Printf("found user: %+v\n", u)
}
