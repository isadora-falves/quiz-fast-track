package repositories

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
	repo := NewMemoryQuizRepository(&quizScores)

	got := repo.GetAllScores()
	g.Expect(*got).To(Equal([]float64{5.3, 7.9, 9.1}))
}	

func TestMemoryQuizRepositorySave(t *testing.T) {
	g := NewGomegaWithT(t)
	initialScores := []entities.QuizScore{}
	repo := NewMemoryQuizRepository(&initialScores)

	newQuizScore := entities.QuizScore{UserName: "New User", Score: 8.5}

	err := repo.Save(newQuizScore)

	g.Expect(err).NotTo(HaveOccurred(), "Save should not produce an error")
	g.Expect(repo.(*memoryQuizRepository).nextID).To(Equal(2))
}
