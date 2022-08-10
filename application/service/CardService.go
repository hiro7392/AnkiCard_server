package service

import (
	"fmt"
	"log"

	"github.com/sakana7392/AnkiCard_server/infra/repository"
	"github.com/sakana7392/AnkiCard_server/testData"
)

// テストデータをDBに挿入
func InsertTestCardData() {
	for _, card := range testData.TestCardData {
		err := repository.CreateNewCard_DB(&card)
		if err != nil {
			fmt.Println("test Data Card insert error")
			log.Println(err)
		}
	}
	fmt.Println("test Card Data inserted successfully")
}
