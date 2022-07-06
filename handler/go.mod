module github.com/sakana7392/AnkiCard_server/handler

go 1.16

require (
	github.com/sakana7392/AnkiCard_server/domain v0.0.0-20220701005301-fe748a022936
	github.com/sakana7392/AnkiCard_server/repository v0.0.0-20220701014141-25b86d4b2eeb
)

replace github.com/sakana7392/AnkiCard_server/repository => ../repository
