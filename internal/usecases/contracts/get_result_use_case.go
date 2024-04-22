package contracts

import (
	"quiz-fast-track/internal/usecases/ports/input"
	"quiz-fast-track/internal/usecases/ports/output"
)

type GetResultUseCase interface {
	Execute(input.AnswersInput) (output.ResultOutput, error)
}