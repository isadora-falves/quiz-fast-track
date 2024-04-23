package main

// @title Quiz Fast Track API
// @version 1
// @description This is a quiz API server.
// @BasePath /api/v1
import (
	"fmt"
	"net/http"
	"quiz-fast-track/internal/entities"
	"quiz-fast-track/internal/infra/database/memory/repositories"
	"quiz-fast-track/internal/usecases"
	"quiz-fast-track/internal/usecases/ports/input"

	_ "quiz-fast-track/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	// Load memory repositories
	questions := []entities.Question{
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

	quizScore := []entities.QuizScore{
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
	
	memoryQuestionsRepository := repositories.NewMemoryQuestionsRepository(&questions)
	memoryQuizRespository := repositories.NewMemoryQuizRepository(&quizScore)

	// Load use cases
	getQuestionsUseCase := usecases.NewGetQuestionsUseCase(memoryQuestionsRepository)
	correctQuizUseCase := usecases.NewCorrectQuizUseCase(memoryQuestionsRepository, memoryQuizRespository)

	fmt.Println(getQuestionsUseCase.Execute())
	fmt.Println(correctQuizUseCase.Execute(input.QuizInput{}))
	r := gin.Default()
	v1 := r.Group("api/v1")

	// Define the endpoint for retrieving questions
	// @Summary Retrieve all questions
	// @Description get questions
	// @Produce json
	// @Success 200 {array} entities::Question
	// @Router /questions [get]
	v1.GET("/questions", func(c *gin.Context) {
		questions := getQuestionsUseCase.Execute()
		c.JSON(http.StatusOK, questions)
	})

		// Define the endpoint for quiz correction
		// @Summary Correct quiz and return score
		// @Description get quiz score
		// @Produce json
		// @Success 200 {object} entities::QuizScore
		// @Router /quiz [get]
		v1.GET("/quiz", func(c *gin.Context) {
			quizOutput, _ := correctQuizUseCase.Execute(input.QuizInput{})

			c.JSON(http.StatusOK, quizOutput)
		})
		
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		r.Run(":3000")
}