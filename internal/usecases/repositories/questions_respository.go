package repositories

import "quiz-fast-track/internal/entities"

type QuestionsRepository interface {
	GetAll() *[]entities.Question
	FindQuestionById(id int) (*entities.Question, error)
}
