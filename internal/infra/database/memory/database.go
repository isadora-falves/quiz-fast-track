package memory

import (
	"quiz-fast-track/internal/entities"
)

func LoadQuestions() []entities.Question {
	return []entities.Question{
		{
			Id:   1,
			Text: "What is the capital of France?",
			Alternatives: []entities.Alternative{
				{
					Id:        1,
					Text:      "Paris",
					IsCorrect: true,
				},
				{
					Id:        2,
					Text:      "London",
					IsCorrect: false,
				},
				{
					Id:        3,
					Text:      "Madrid",
					IsCorrect: false,
				},
			},
		},
		{
			Id:   2,
			Text: "What is the capital of Spain?",
			Alternatives: []entities.Alternative{
				{
					Id:        1,
					Text:      "Paris",
					IsCorrect: false,
				},
				{
					Id:        2,
					Text:      "London",
					IsCorrect: false,
				},
				{
					Id:        3,
					Text:      "Madrid",
					IsCorrect: true,
				},
			},
		},
	}
}

func LoadQuizScores() []entities.QuizScore {
	return []entities.QuizScore{
		{
			Id:       1,
			UserName: "John Doe",
			Score:    1.5,
		},
		{
			Id:       2,
			UserName: "Jane Doe",
			Score:    7.5,
		},
		{
			Id:       3,
			UserName: "Alice",
			Score:    9.5,
		},
	}
}
