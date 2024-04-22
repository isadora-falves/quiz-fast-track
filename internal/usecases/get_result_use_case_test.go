package usecases

import (
	"quiz-fast-track/internal/usecases/mocks"
	"testing"

	. "github.com/onsi/gomega"
)

func TestReturnCorrectResponse(t *testing.T){
	// arrange
	g := NewWithT(t)

	questionsRepository := mocks.NewQuestionsRepository(t)
	questionsRepository.On("GetAnswerById", 1).Return("B", nil)
	questionsRepository.On("GetAnswerById", 2).Return("A", nil)
	getResultUseCase := GetResultUseCase(questionsRepository)

	expectedResponse := ResultOutput{	
		Resume: "You answered 1 questions correctly out of 2. You made 1 errors.",
		Hits: 2,
		Erros: 1,
		QuizTemplate []QuizTemplateOutput
	}}

	// act
	response := getResultUseCase.Execute()

	// assert
	g.Expect(response).To(Equal(expectedResponse))
}