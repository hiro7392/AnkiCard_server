package testData

import "github.com/sakana7392/AnkiCard_server/domain/model"

var TestTagData []model.Tag = []model.Tag{
	{
		TagId:         1,
		CreatedUserId: 2,
		TagName:       "TCP/IP",
	},
	{
		TagId:         2,
		CreatedUserId: 3,
		TagName:       "データベース",
	},
	{
		TagId:         3,
		CreatedUserId: 4,
		TagName:       "英語",
	},
	{
		TagId:         4,
		CreatedUserId: 2,
		TagName:       "アルゴリズム",
	},
}
