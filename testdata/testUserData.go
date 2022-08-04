package testData

import "github.com/sakana7392/AnkiCard_server/domain/model"

var TestUserData [3]model.User = [3]model.User{
	{
		UserId:         1,
		UserName:       "sakana",
		NextQuestionId: 3,
		UserLevel:      1,
		Email:          "hiromiimkw@gmail.com",
		Password:       "password",
	},
	{
		UserId:         2,
		UserName:       "sakana2",
		NextQuestionId: 5,
		UserLevel:      2,
		Email:          "sakana2w@gmail.com",
		Password:       "password2",
	},
	{
		UserId:         3,
		UserName:       "sakana3",
		NextQuestionId: 4,
		UserLevel:      1,
		Email:          "sakana3@gmail.com",
		Password:       "password3",
	},
}
