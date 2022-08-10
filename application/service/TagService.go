package service

import (
	"fmt"
	"log"

	"github.com/sakana7392/AnkiCard_server/infra/repository"
	"github.com/sakana7392/AnkiCard_server/testData"
)

// テストデータをDBに挿入
func InsertTestTagData() {
	for _, tag := range testData.TestTagData {
		err := repository.CreateNewTag_DB(&tag)
		if err != nil {
			fmt.Println("test Data Card insert error")
			log.Println(err)
		}
	}
	fmt.Println("test tag Data inserted successfully")
}
