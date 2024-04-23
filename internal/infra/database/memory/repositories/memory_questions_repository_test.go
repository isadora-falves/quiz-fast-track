package repositories

import (
	"errors"
	"quiz-fast-track/internal/entities"
	"testing"

	. "github.com/onsi/gomega"
)

func TestMemoryQuestionsRepositoryGetAll(t *testing.T) {
	g := NewGomegaWithT(t)
	questions := []entities.Question{
		{Id: 1, Text: "Question 1"},
		{Id: 2, Text: "Question 2"},
		{Id: 3, Text: "Question 3"},
	}

	repo := NewMemoryQuestionsRepository(&questions)

	got := repo.GetAll()
	g.Expect(*got).To(Equal(questions))
}

func TestMemoryQuestionsRepositoryFindQuestionById(t *testing.T) {
	g := NewGomegaWithT(t)
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

