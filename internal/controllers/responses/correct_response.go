package responses

type CorrectResponse struct {
	Resume       string         `json:"resume" example:"You answered 1 question correctly out of 2. You made 1 error. You were better than 20% of all quizzers"`
	RightAnswers int            `json:"right_answers" example:"1"`
	WrongAnswers int            `json:"wrong_answers" example:"1"`
	QuizTemplate []QuizResponse `json:"quizes"`
}
