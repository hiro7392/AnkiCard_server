package model

type Card struct {
	CardId        int
	TagId         int
	CreatedUserId int
	LearningLevel int
	QuestionText  string
	AnswerText    string
}