# GoMint / Postgres

## Как подключаться?
Через `pgx`-драйвер и `database/sql`:
```go
db, _ := postgres.Open(ctx, postgres.Config{ Host:"127.0.0.1", Port:15432, User:"gomint", Password:"gomint", DBName:"gomint", SSLMode:"disable" })
