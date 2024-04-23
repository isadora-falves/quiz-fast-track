package memory

import (
	"errors"
	"quiz-fast-track/internal/entities"
	"quiz-fast-track/internal/usecases/repositories"
)

type memoryQuestionsRepository struct {
	questions map[int]entities.Question
}

func NewMemoryQuestionsRepository() repositories.QuestionsRepository {
	return &memoryQuestionsRepository{
		questions: make(map[int]entities.Question),
	}
}

func (m *memoryQuestionsRepository) GetAll() *[]entities.Question {
	questions := make([]entities.Question, 0, len(m.questions))
	for _, question := range m.questions {
			questions = append(questions, question)
	}
	return &questions
}

func (m *memoryQuestionsRepository) FindQuestionById(id int) (*entities.Question, error) {
	question, exists := m.questions[id]
	if !exists {
		return nil, errors.New("question not found")
	}
	return &question, nil
}

func (m *memoryQuestionsRepository) SaveQuestion(question entities.Question) error {
	m.questions[question.Id] = question
	return nil
}