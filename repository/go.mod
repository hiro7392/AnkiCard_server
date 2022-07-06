module github.com/sakana7392/AnkiCard_server/repository

go 1.16

require (
	github.com/sakana7392/AnkiCard_server/domain v0.0.0-20220701005301-fe748a022936
	github.com/sakana7392/AnkiCard_server/infra v0.0.0-20220701005301-fe748a022936
)

replace github.com/sakana7392/AnkiCard_server/domain => ../domain

replace github.com/sakana7392/AnkiCard_server/infra => ../infra
