package infra

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql" //コード内で直接参照するわけではないが、依存関係のあるパッケージには最初にアンダースコア_をつける
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", "root@/AnkiCard?parseTime=true")

	if err != nil {
		panic(err.Error())
	}

	log.Println("init success!\n Connected to mysql")
}
