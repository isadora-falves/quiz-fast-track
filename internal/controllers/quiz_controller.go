package controllers

import (
	"fmt"
	"net/http"
	"quiz-fast-track/internal/controllers/requests"
	"quiz-fast-track/internal/controllers/responses"
	"quiz-fast-track/internal/usecases/contracts"
	"quiz-fast-track/internal/usecases/ports/input"
	"quiz-fast-track/internal/usecases/ports/output"

	"github.com/gin-gonic/gin"
)

type quizController struct {
	correctQuizUseCase  contracts.CorrectQuizUseCase
	getQuestionsUseCase contracts.GetQuestionsUseCase
}

func NewQuizController(
	correctQuizUseCase contracts.CorrectQuizUseCase,
	getQuestionsUseCase contracts.GetQuestionsUseCase,
) *quizController {
	return &quizController{
		correctQuizUseCase:  correctQuizUseCase,
		getQuestionsUseCase: getQuestionsUseCase,
	}
}

// @Summary Retrieve all questions
// @Description get questions
// @Produce json
// @Success 200 {array} responses.QuestionResponse
// @Router /questions [get]
func (qc *quizController) Get(ctx *gin.Context) {
	questions := qc.getQuestionsUseCase.Execute()

	var response []responses.QuestionResponse

	for _, question := range questions {
		var alternativeResponses []responses.AlternativeResponse

		for _, alternative := range question.Alternatives {
			alternativeResponses = append(alternativeResponses, responses.AlternativeResponse{
				Id:     alternative.Id,
				Text:   alternative.Text,
				Option: alternative.Option,
			})
		}

		response = append(response, responses.QuestionResponse{
			Id:           question.Id,
			Text:         question.Text,
			Alternatives: alternativeResponses,
		})
	}

	ctx.JSON(http.StatusOK, response)
}

// @Summary Correct a quiz
// @Description correct a quiz
// @Produce json
// @Router /quiz [post]
// @Param correct body requests.CorrectRequest true "Correct Request"
// @Success 200 {object} responses.CorrectResponse
func (qc *quizController) Correct(ctx *gin.Context) {
	var request requests.CorrectRequest
	fmt.Println(request)
	fmt.Println(ctx.BindJSON(&request))
	fmt.Println(ctx.Params)

	convertedAnswers := ConvertAnswerRequestsToAnswerInputs(request.Answers)
	quizInput := input.QuizInput{
		User:    request.User,
		Answers: convertedAnswers,
	}
	
	correctedQuiz, err := qc.correctQuizUseCase.Execute(quizInput)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	quizTemplateResponses := ConverQuizTemplatetResponse(correctedQuiz.QuizTemplate)

	response := responses.CorrectResponse{
		Resume:       correctedQuiz.Resume,
		RightAnswers: correctedQuiz.RightAnswers,
		WrongAnswers: correctedQuiz.WrongAnswers,
		QuizTemplate: quizTemplateResponses,
	}

	// Send the response
	ctx.JSON(http.StatusOK, response)
}

func ConvertAnswerRequestsToAnswerInputs(requests []requests.AnswerRequest) []input.AnswerInput {
	var inputs []input.AnswerInput
	for _, req := range requests {
		input := input.AnswerInput{
			QuestionId:     req.QuestionId,
			Option: req.Option,
		}
		inputs = append(inputs, input)
	}
	return inputs
}


func ConverQuizTemplatetResponse(outputs []output.QuizTemplateOutput) []responses.QuizResponse {
	var quizResponses []responses.QuizResponse
	for _, out := range outputs {
			quizResponse := responses.QuizResponse{
					QuestionId:     out.QuestionId,
					SelectedOption: out.SelectedOption,
					CorrectOption:  out.CorrectOption,
			}
			quizResponses = append(quizResponses, quizResponse)
	}
	return quizResponses
}