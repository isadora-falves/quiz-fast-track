package repositories

import "quiz-fast-track/internal/entities"

type QuestionsRepository interface {
	GetAll() (*[]entities.Question)
	GetAnswerById(id int) (*entities.Alternative, error)
}
