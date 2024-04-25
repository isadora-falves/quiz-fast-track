package repositories

import (
	"errors"
	"quiz-fast-track/internal/entities"
	"testing"

	. "github.com/onsi/gomega"
)

func TestMemoryQuestionsRepositoryGetAll(t *testing.T) {
	// Arrange
	g := NewGomegaWithT(t)
	questions := []entities.Question{
		{Id: 1, Text: "Question 1"},
		{Id: 2, Text: "Question 2"},
		{Id: 3, Text: "Question 3"},
	}

	repo := NewMemoryQuestionsRepository(&questions)

	// Act
	response := repo.GetAll()

	// Assert
	g.Expect(*response).To(Equal(questions))
}

func TestFindQuestionByIdReturnsOneQuestion(t *testing.T) {
	// Arrange
	g := NewWithT(t)
	questions := []entities.Question{
		{Id: 1, Text: "Question 1"},
		{Id: 2, Text: "Question 2"},
	}
	repo := NewMemoryQuestionsRepository(&questions)

	expectedResponse := &entities.Question{
		Id:   2,
		Text: "Question 2",
	}

	// Act
	response, err := repo.FindQuestionById(2)

	// Assert
	g.Expect(err).To(BeNil())
	g.Expect(response).To(Equal(expectedResponse))
}

func TestFindQuestionByIdReturnsError(t *testing.T) {
	// Arrange
	g := NewWithT(t)
	questions := []entities.Question{
		{Id: 1, Text: "Question 1"},
		{Id: 2, Text: "Question 2"},
	}
	repo := NewMemoryQuestionsRepository(&questions)

	// Act
	got, err := repo.FindQuestionById(3)

	// Assert
	g.Expect(err).To(MatchError(errors.New("question not found")))
	g.Expect(got).To(BeNil())
}
