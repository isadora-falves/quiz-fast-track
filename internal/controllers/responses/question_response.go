package responses

type QuestionResponse struct {
	Id           int                   `json:"id" example:"1"`
	Text         string                `json:"text" example:"What is the best country in the world?"`
	Alternatives []AlternativeResponse `json:"alternatives"`
}
