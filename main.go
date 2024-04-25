package main

import (
	"quiz-fast-track/internal/controllers"
	"quiz-fast-track/internal/infra/database/memory"
	"quiz-fast-track/internal/infra/database/memory/repositories"
	"quiz-fast-track/internal/usecases"

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

	quizController := controllers.NewQuizController(correctQuizUseCase, getQuestionsUseCase)

	r := gin.Default()
	v1 := r.Group("api/v1")

	v1.GET("/questions", quizController.Get)

	v1.POST("/quiz", quizController.Correct)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":3000")
}
