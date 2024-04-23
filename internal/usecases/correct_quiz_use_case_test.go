package usecases

import (
	"quiz-fast-track/internal/usecases/mocks"
	"quiz-fast-track/internal/usecases/ports/input"
	"quiz-fast-track/internal/usecases/ports/output"
	"testing"

	. "github.com/onsi/gomega"
)

func TestReturnCorrectResponse(t *testing.T) {
	// arrange
	g := NewWithT(t)

	questions := getQuestions()
	quizScores := &[]float64{0.5, 0.36, 0.66, 0.75, 0.9}

	questionsRepository := mocks.NewQuestionsRepository(t)
	questionsRepository.On("FindQuestionById", 1).Return(&questions[0], nil)
	questionsRepository.On("FindQuestionById", 2).Return(&questions[1], nil)
	
	quizRepository := mocks.NewQuizRepository(t)
	quizRepository.On("GetAllScores").Return(quizScores, nil)

	input := input.QuizInput{
		User: "Isadora Alves",
		Answers: []input.AnswerInput{
			{
				QuestionId: 1,
				Option:     "B",
			},
			{
				QuestionId: 2,
				Option:     "B",
			},
		},
	}

	expectedResponse := getCorrectQuizResponse()
	correctQuizUseCase := NewCorrectQuizUseCase(questionsRepository, quizRepository)

	// act
	response, err := correctQuizUseCase.Execute(input)

	// assert
	g.Expect(response).To(Equal(expectedResponse))
	g.Expect(err).To(BeNil())
}

func TestReturnWhenWeHaveDuplicateAnswers(t *testing.T) {
	// arrange
	g := NewWithT(t)

	questions := getQuestions()

	questionsRepository := mocks.NewQuestionsRepository(t)
	questionsRepository.On("FindQuestionById", 1).Return(&questions[0], nil)
	questionsRepository.On("FindQuestionById", 1).Return(&questions[1], nil)

	quizRepository := mocks.NewQuizRepository(t)

	input := input.QuizInput{
		User: "Isadora Alves",
		Answers: []input.AnswerInput{
			{
				QuestionId: 1,
				Option:     "B",
			},
			{
				QuestionId: 1,
				Option:     "A",
			},
		},
	}

	correctQuizUseCase := NewCorrectQuizUseCase(questionsRepository, quizRepository)

	// act
	response, err := correctQuizUseCase.Execute(input)

	// assert
	g.Expect(response).To(BeNil())
	g.Expect(err).ToNot(BeNil())
}


func getCorrectQuizResponse() *output.QuizOutput {
	return &output.QuizOutput{
		Resume:       "You answered 1 question correctly out of 2. You made 1 error. You were better than 20% of all quizzers",
		RightAnswers: 1,
		WrongAnswers: 1,
		QuizTemplate: []output.QuizTemplateOutput{
			{
				QuestionId:     1,
				SelectedOption: "B",
				CorrectOption:  "B",
			},
			{
				QuestionId:     2,
				SelectedOption: "B",
				CorrectOption:  "A",
			},
		},
	}
}
