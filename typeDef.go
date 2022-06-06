type Card struct {
	CardId        int
	TagId         int
	LearinigLevel int
	QuestionText  string
	AnswertText   string
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