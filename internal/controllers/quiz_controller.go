package controllers

import (
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
// @Param correct body requests.QuizRequest true "Correct Request"
// @Success 200 {object} responses.QuizResponse
func (qc *quizController) Correct(ctx *gin.Context) {
	var request requests.QuizRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	input := qc.createQuizInput(request)

	output, err := qc.correctQuizUseCase.Execute(input)

	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	response := qc.createQuizSuccessResponse(output)

	ctx.JSON(http.StatusOK, response)
}

func (qc *quizController) createQuizInput(request requests.QuizRequest) input.QuizInput {
	var answers []input.AnswerInput

	for _, req := range request.Answers {
		answers = append(answers, input.AnswerInput{
			QuestionId: req.QuestionId,
			Option:     req.Option,
		})
	}

	return input.QuizInput{
		User:    request.User,
		Answers: answers,
	}
}

func (qc *quizController) createQuizSuccessResponse(output *output.QuizOutput) responses.QuizResponse {
	var quizResponses []responses.QuizTemplateResponse
	for _, template := range output.QuizTemplate {
		quizTemplate := responses.QuizTemplateResponse{
			QuestionId:     template.QuestionId,
			SelectedOption: template.SelectedOption,
			CorrectOption:  template.CorrectOption,
		}
		quizResponses = append(quizResponses, quizTemplate)
	}

	return responses.QuizResponse{
		Resume:       output.Resume,
		RightAnswers: output.RightAnswers,
		WrongAnswers: output.WrongAnswers,
		QuizTemplate: quizResponses,
	}
}
