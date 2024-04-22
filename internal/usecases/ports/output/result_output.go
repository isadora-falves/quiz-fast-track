package output

type ResultOutput struct {
	Resume string
	Hits int
	Erros int
	QuizTemplate []QuizTemplateOutput
}
