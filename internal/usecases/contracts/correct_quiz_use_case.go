package contracts

import (
	"quiz-fast-track/internal/usecases/ports/input"
	"quiz-fast-track/internal/usecases/ports/output"
)

type CorrectQuizUseCase interface {
	Execute(input.QuizInput) (*output.QuizOutput, error)
}
