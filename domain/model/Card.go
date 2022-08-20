package model

type Card struct {
	CardId        int
	TagId         int
	TagName       string
	CreatedUserId int
	LearningLevel int
	QuestionText  string
	AnswerText    string
}