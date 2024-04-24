package repositories

import "quiz-fast-track/internal/entities"

type QuizRepository interface {
	GetAllScores() *[]float64
	Save(quizScore entities.QuizScore) error
}
