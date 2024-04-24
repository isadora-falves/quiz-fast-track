package contracts

import "quiz-fast-track/internal/usecases/ports/output"

type GetQuestionsUseCase interface {
	Execute() []output.QuestionOutput
}
