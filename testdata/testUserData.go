package testData

import "github.com/sakana7392/AnkiCard_server/domain/model"

var TestUserData []model.User = []model.User{
	{
		UserId:         1,
		UserName:       "sakana",
		NextQuestionId: 3,
		UserLevel:      1,
		Email:          "hiromiimkw@gmail.com",
		Password:       "pass",
	},
	{
		UserId:         2,
		UserName:       "sakana2",
		NextQuestionId: 5,
		UserLevel:      2,
		Email:          "sakanatwo@gmail.com",
		Password:       "password2",
	},
	{
		UserId:         3,
		UserName:       "sakana3",
		NextQuestionId: 4,
		UserLevel:      1,
		Email:          "sakanathree@gmail.com",
		Password:       "password3",
	},
	{
		UserId:         4,
		UserName:       "sakana4",
		NextQuestionId: 4,
		UserLevel:      4,
		Email:          "sakanafour@gmail.com",
		Password:       "password4",
	},
}
