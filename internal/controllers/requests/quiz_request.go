package requests

type QuizRequest struct {
	User    string          `json:"user" example:"John Doe"`
	Answers []AnswerRequest `json:"answers"`
}
