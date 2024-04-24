package requests

type CorrectRequest struct {
	User 	string `json:"user" example:"Isadora Alves"`
	Answers []AnswerRequest `json:"answers"`
}