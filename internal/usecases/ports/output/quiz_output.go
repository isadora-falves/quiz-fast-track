package output

type QuizOutput struct {
	Resume       string
	RightAnswers int
	WrongAnswers int
	QuizTemplate []QuizTemplateOutput
}