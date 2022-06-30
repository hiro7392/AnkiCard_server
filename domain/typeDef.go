package domain

type Card struct {
	CardId        int
	TagId         int
	LearningLevel int
	QuestionText  string
	AnswerText    string
}
type User struct {
	UserId         int
	UserName       int
	NextQuestionId int
	UserLevel      int
}
type Tag struct {
	TagId   int
	TagName string
}
