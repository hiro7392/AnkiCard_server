//module AnkiCard_server
module github.com/sakana7392/AnkiCard_server

go 1.16

//replace github.com/sakana7392/AnkiCard_server/handler => ./handler
require (
	github.com/sakana7392/AnkiCard_server/domain v0.0.0-20220701005301-fe748a022936
	github.com/sakana7392/AnkiCard_server/infra v0.0.0-20220701005301-fe748a022936
)
