package memory

import (
	"quiz-fast-track/internal/entities"
	"quiz-fast-track/internal/usecases/repositories"
)

type memoryQuizRepository struct {}

func NewMemoryQuizRepository() repositories.QuizRepository {
	return &memoryQuizRepository{}
}

func (m *memoryQuizRepository) GetAllScores() *[]float64 {
	return nil
}

func (m *memoryQuizRepository) Save(quizScore entities.QuizScore) error {
	return nil
}

