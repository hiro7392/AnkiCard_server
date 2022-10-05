package infra

import (
	"database/sql"
	"log"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql" //コード内で直接参照するわけではないが、依存関係のあるパッケージには最初にアンダースコア_をつける
)

var Db *sql.DB

type NullableCard struct {
	CardId        int
	TagId         sql.NullInt64
	TagName       sql.NullString
	CreatedUserId int
	LearningLevel int
	QuestionText  string
	AnswerText    string
}

func init() {
	var err error
	//Db, err = sql.Open("mysql", "root:mysql@/AnkiCard?parseTime=true")
	jst, err := time.LoadLocation("Asia/Tokyo")
	c := mysql.Config{
		DBName:    os.Getenv("DB_NAME"),
		User:      os.Getenv("DB_USER"),
		Passwd:    os.Getenv("DB_PASSWORD"),
		Addr:      os.Getenv("DB_ADDR"),
		Net:       "tcp",
		ParseTime: true,
		Collation: "utf8mb4_unicode_ci",
		Loc:       jst,
	}
	Db, err = sql.Open("mysql", c.FormatDSN())

	if err != nil {
		panic(err.Error())
	}
	log.Println("init success!\n Connected to mysql")
}
