package repositories

import (
	"errors"
	"quiz-fast-track/internal/entities"
	"quiz-fast-track/internal/usecases/repositories"
)

type memoryQuestionsRepository struct {
	questions *[]entities.Question
}

func NewMemoryQuestionsRepository(questions *[]entities.Question) repositories.QuestionsRepository {
	return &memoryQuestionsRepository{
		questions: questions,
	}
}

// GetAll returns all questions
func (m *memoryQuestionsRepository) GetAll() *[]entities.Question {
	return m.questions
}

// FindQuestionById returns a question by id
func (m *memoryQuestionsRepository) FindQuestionById(id int) (*entities.Question, error) {
	for _, question := range *m.questions {
		if question.Id == id {
			return &question, nil
		}
	}
	return nil, errors.New("question not found")
}
