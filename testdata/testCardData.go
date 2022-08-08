package testData

import "github.com/sakana7392/AnkiCard_server/domain/model"

var TestCardData []model.Card = []model.Card{
	{
		CardId:        1,
		TagId:         1,
		CreatedUserId: 1,
		LearningLevel: 1,
		QuestionText:  "サンプル問題 TCP/IPプロトコルにおける 3way hand shakeとは?",
		AnswerText:    "TCPはコネクション型プロトコル（通信相手の応答があってはじめて通信を開始する）であることから、データ転送を行う前にコネクションの確立を行います。このTCPにおいて使用されるコネクションの確立のことを3ウェイハンドシェイクといいます。以下の手順の通り3回のやりとりによって確立されます",
	},
	{
		CardId:        2,
		TagId:         2,
		CreatedUserId: 1,
		LearningLevel: 1,
		QuestionText:  "サンプル問題 DBのACID特性とは？",
		AnswerText:    "ACID特性とは、トランザクション処理において必要とされる4つの要素、Atomicity（原子性）、Consistency（一貫性）、Isolation（独立性）、Durability（永続性）を頭字語で表したものである",
	},
	{
		CardId:        3,
		TagId:         2,
		CreatedUserId: 3,
		LearningLevel: 2,
		QuestionText:  "外部キー制約とは？",
		AnswerText:    "RDBにおいて、他のテーブルのカラムと対応している場合に、そのカラムに存在しない要素(Idなど)を持つレコードを挿入することができないこと。その逆も存在する",
	},
}
