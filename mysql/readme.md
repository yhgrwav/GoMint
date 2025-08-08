# GoMint / MySQL

Подключение: драйвер github.com/go-sql-driver/mysql, плейсхолдеры — знак вопроса (?).
Отличия от Postgres: placeholders ?, LastInsertId() работает из коробки, строковые типы лучше держать в utf8mb4.

Быстрый старт:
1) Поднять сервис:
   cd docker
   docker compose up -d

2) Запустить пример:
   cd ..
   go run ./examples/mysql_ping
