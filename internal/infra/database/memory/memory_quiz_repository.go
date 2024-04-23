package memory

import (
	"quiz-fast-track/internal/entities"
	"quiz-fast-track/internal/usecases/repositories"
)

type memoryQuizRepository struct {
	quizScores *[]entities.QuizScore
	nextID     int
}

func NewMemoryQuizRepository(quizScores *[]entities.QuizScore) repositories.QuizRepository {
	return &memoryQuizRepository{
		quizScores: quizScores,
		nextID:     1,
	}
}

func (m *memoryQuizRepository) GetAllScores() *[]float64 {
	var scores []float64
	for _, quizScore := range *m.quizScores {
		scores = append(scores, quizScore.Score)
	}
	return &scores
}

func (m *memoryQuizRepository) Save(quizScore entities.QuizScore) error {
	quizScore.Id = m.nextID
	m.nextID++
	*m.quizScores = append(*m.quizScores, quizScore)
	return nil
}

