package memory

import (
	"quiz-fast-track/internal/entities"
	"testing"

	. "github.com/onsi/gomega"
)

func TestMemoryQuizRepositoryGetAll(t *testing.T) {
	g := NewWithT(t)
	quizScores := []entities.QuizScore{
		{Id: 1, UserName: "John Doe", Score: 5.3,},
		{Id: 2, UserName: "Jane Doe", Score: 7.9,},
		{Id:3, UserName: "John Smith", Score: 9.1,},
	}
	repo := NewMemoryQuizRepository()
	
	repo.(*memoryQuizRepository).quizScores = make(map[int]entities.QuizScore)

	for _, quizScore := range quizScores {
		repo.(*memoryQuizRepository).quizScores[quizScore.Id] = quizScore
	}

	expectedScores := []float64{5.3, 7.9, 9.1}

	got := repo.GetAllScores()
	g.Expect(got).To(Equal(&expectedScores))
}	

func TestMemoryQuizRepositorySave(t *testing.T) {
	g := NewWithT(t)
	quizScore := entities.QuizScore{Id: 1,UserName: "John Doe", Score: 5.3}
	repo := NewMemoryQuizRepository()
	
	err := repo.Save(quizScore)
	g.Expect(err).To(BeNil())
	g.Expect(repo.(*memoryQuizRepository).quizScores[1]).To(Equal(quizScore))
}
