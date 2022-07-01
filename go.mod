module github.com/sakana7392/AnkiCard_server/main

go 1.16

replace github.com/sakana7392/AnkiCard_server/handler => ./handler

require (
	github.com/go-sql-driver/mysql v1.6.0
	github.com/sakana7392/AnkiCard_server/domain v0.0.0-20220701005301-fe748a022936
	github.com/sakana7392/AnkiCard_server/handler v0.0.0-00010101000000-000000000000
	github.com/stretchr/testify v1.8.0
)

replace github.com/sakana7392/AnkiCard_server/test => ./test
