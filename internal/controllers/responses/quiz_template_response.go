package responses

type QuizTemplateResponse struct {
	QuestionId     int    `json:"question_id" example:"1"`
	SelectedOption string `json:"selected_option" example:"A"`
	CorrectOption  string `json:"correct_option" example:"A"`
}
