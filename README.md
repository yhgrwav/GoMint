# GoMint

GoMint — a collection of ready-to-use Go code snippets for working with databases and APIs.  
Focus: PostgreSQL, MySQL, REST client, Docker setups.

## Modules
- postgres — connections, queries, transactions (PostgreSQL)
- mysql — connections, queries, transactions (MySQL)
- rest — JSON client (GET/POST, timeouts, basic retries)
- docker — compose files for local DBs
- examples — runnable mini demos

---

## Quick start (PostgreSQL)
1. Start local DB:
   cd docker
   docker compose up -d

2. Run example:
   cd ..
   go run ./examples/postgres_ping

---

## REST examples
Simple JSON client with timeouts and basic retries (429/5xx).

GET:
go run ./examples/rest_get

POST:
go run ./examples/rest_post

Code lives in:
- rest/client.go — client implementation
- examples/rest_get/main.go — GET demo
- examples/rest_post/main.go — POST demo

---

## Why $1, $2 in Postgres?
PostgreSQL uses positional parameters ($1, $2, ...).  
It’s safe (no SQL injection via values) and can reuse query plans.

## License
MIT

## Install
go get github.com/yhgrwav/GoMint/postgres
go get github.com/yhgrwav/GoMint/mysql
go get github.com/yhgrwav/GoMint/rest

## Import
import "github.com/yhgrwav/GoMint/postgres"
import "github.com/yhgrwav/GoMint/mysql"
import "github.com/yhgrwav/GoMint/rest"
