package usecases

import (
	"quiz-fast-track/internal/entities"
	"quiz-fast-track/internal/usecases/mocks"
	"quiz-fast-track/internal/usecases/ports/output"

	. "github.com/onsi/gomega"

	"testing"
)

func TestExecuteShouldReturnTwoQuestions(t *testing.T) {
	// arrange
	g := NewWithT(t)

	questions := getQuestions()
	questionsRepository := mocks.NewQuestionsRepository(t)
	questionsRepository.On("GetAll").Return(&questions)

	expectedResponse := getTwoQuestionsOutput()
	getQuestionsUseCase := GetQuestionsUseCase(questionsRepository)

	// act
	response := getQuestionsUseCase.Execute()

	// assert
	g.Expect(response).To(Equal(expectedResponse))
}

func getQuestions() []entities.Question {
	return []entities.Question{
		{
			Id:   1,
			Text: "What is the most abundant chemical element in the universe?",
			Alternatives: []entities.Alternative{
				{
					Id:        1,
					Option:    "A",
					Text:      "Oxygen",
					IsCorrect: false,
				},
				{
					Id:        2,
					Option:    "B",
					Text:      "Hydrogen",
					IsCorrect: true,
				},
			},
		},
		{
			Id:   2,
			Text: "What was the first computer program to defeat a reigning world chess champion under tournament conditions?",
			Alternatives: []entities.Alternative{
				{
					Id:        3,
					Option:    "A",
					Text:      "Deep Blue",
					IsCorrect: true,
				},
				{
					Id:        4,
					Option:    "B",
					Text:      "AlphaGo",
					IsCorrect: false,
				},
			},
		},
	}
}

func getTwoQuestionsOutput() []output.QuestionOutput {
	return []output.QuestionOutput{
		{
			Id:   1,
			Text: "What is the most abundant chemical element in the universe?",
			Alternatives: []output.AlternativeOutput{
				{
					Id:     1,
					Option: "A",
					Text:   "Oxygen",
				},
				{
					Id:     2,
					Option: "B",
					Text:   "Hydrogen",
				},
			},
		},
		{
			Id:   2,
			Text: "What was the first computer program to defeat a reigning world chess champion under tournament conditions?",
			Alternatives: []output.AlternativeOutput{
				{
					Id:     3,
					Option: "A",
					Text:   "Deep Blue",
				},
				{
					Id:     4,
					Option: "B",
					Text:   "AlphaGo",
				},
			},
		},
	}
}
