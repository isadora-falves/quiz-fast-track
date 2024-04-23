package memory

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
	questions := []entities.Question{
			{Id: 1, Text: "Question 1"},
			{Id: 2, Text: "Question 2"},
			{Id: 3, Text: "Question 3"},
	}

	repo := NewMemoryQuestionsRepository(&questions)

	tests := []struct {
			name    string
			id      int
			want    *entities.Question
			wantErr error
	}{
			{"Existing question", 2, &questions[1], nil},
			{"Non-existing question", 4, nil, errors.New("question not found")},
	}

	for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
					g := NewGomegaWithT(t)
					got, err := repo.FindQuestionById(tt.id)
					g.Expect(got).To(Equal(tt.want), "Check fetched question")
					if tt.wantErr == nil {
							g.Expect(err).NotTo(HaveOccurred(), "Expected no error")
					} else {
							g.Expect(err).To(MatchError(tt.wantErr), "Check error message")
					}
			})
	}
}

