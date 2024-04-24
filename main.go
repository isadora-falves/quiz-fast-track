package main

import (
	"net/http"
	"quiz-fast-track/internal/infra/database/memory"
	"quiz-fast-track/internal/infra/database/memory/repositories"
	"quiz-fast-track/internal/usecases"
	"quiz-fast-track/internal/usecases/ports/input"

	_ "quiz-fast-track/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Quiz Fast Track API
// @version 1
// @description This is a quiz API server.
// @BasePath /api/v1
func main() {
	// Load memory repositories
	questions := memory.LoadQuestions()
	quizScores := memory.LoadQuizScores()

	memoryQuestionsRepository := repositories.NewMemoryQuestionsRepository(&questions)
	memoryQuizRepository := repositories.NewMemoryQuizRepository(&quizScores)

	// Load use cases
	getQuestionsUseCase := usecases.NewGetQuestionsUseCase(memoryQuestionsRepository)
	correctQuizUseCase := usecases.NewCorrectQuizUseCase(memoryQuestionsRepository, memoryQuizRepository)

	r := gin.Default()
	v1 := r.Group("api/v1")

	// @Summary Retrieve all questions
	// @Description get questions
	// @Produce json
	// @Success 200 {array} entities::Question
	// @Router /questions [get]
	v1.GET("/questions", func(c *gin.Context) {
		questions := getQuestionsUseCase.Execute()
		c.JSON(http.StatusOK, questions)
	})

	// @Summary Correct quiz and return score
	// @Description get quiz score
	// @Produce json
	// @Success 200 {object} entities::QuizScore
	// @Router /quiz [get]
	v1.POST("/quiz", func(c *gin.Context) {
		quizOutput, _ := correctQuizUseCase.Execute(input.QuizInput{})
		c.JSON(http.StatusOK, quizOutput)
	})

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":3000")
}
