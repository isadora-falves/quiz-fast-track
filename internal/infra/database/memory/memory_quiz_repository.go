package memory

import (
	"quiz-fast-track/internal/entities"
	"quiz-fast-track/internal/usecases/repositories"
)

type memoryQuizRepository struct {
	quizScores map[int]entities.QuizScore
	nextID     int
}

func NewMemoryQuizRepository() repositories.QuizRepository {
	return &memoryQuizRepository{
		quizScores: make(map[int]entities.QuizScore),
		nextID:     1,
	}
}

func (m *memoryQuizRepository) GetAllScores() *[]float64 {
	quizScores := make([]float64, 0, len(m.quizScores))
	for _, quizScore := range m.quizScores {
			quizScores = append(quizScores, quizScore.Score)
	}
	return &quizScores
}

func (m *memoryQuizRepository) Save(quizScore entities.QuizScore) error {
	quizScore.Id = m.nextID
	m.quizScores[m.nextID] = quizScore
	m.nextID++  
	return nil
}

