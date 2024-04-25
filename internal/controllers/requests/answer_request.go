package requests

type AnswerRequest struct {
	QuestionId int    `json:"question_id" example:"1"`
	Option     string `json:"option" example:"A"`
}
