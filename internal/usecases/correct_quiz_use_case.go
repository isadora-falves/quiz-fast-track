package usecases

import (
	"fmt"
	"quiz-fast-track/internal/entities"
	"quiz-fast-track/internal/usecases/contracts"
	"quiz-fast-track/internal/usecases/ports/input"
	"quiz-fast-track/internal/usecases/ports/output"
	"quiz-fast-track/internal/usecases/repositories"
	"sort"
)

type correctQuizUseCase struct {
	questionsRepository repositories.QuestionsRepository
	quizRepository      repositories.QuizRepository
}

func NewCorrectQuizUseCase(
	questionsRepository repositories.QuestionsRepository,
	quizRepository repositories.QuizRepository,
) contracts.CorrectQuizUseCase {
	return &correctQuizUseCase{
		questionsRepository: questionsRepository,
		quizRepository:      quizRepository,
	}
}

// Execute corrects the quiz and returns the result
func (uc *correctQuizUseCase) Execute(input input.QuizInput) (*output.QuizOutput, error) {
	questionIDs := make(map[int]bool)
	rightsCount := 0
	wrongsCount := 0
	var quizTemplates []output.QuizTemplateOutput

	for _, answer := range input.Answers {
		if _, exists := questionIDs[answer.QuestionId]; exists {
			return nil, fmt.Errorf("duplicated question ID: %d", answer.QuestionId)
		}

		questionIDs[answer.QuestionId] = true

		question, err := uc.questionsRepository.FindQuestionById(answer.QuestionId)

		if err != nil {
			return nil, err
		}

		correctAlternative, err := question.GetCorrectAlternative()

		if err != nil {
			return nil, err
		}

		if answer.Option == correctAlternative.Option {
			rightsCount++
		} else {
			wrongsCount++
		}

		quizTemplates = append(quizTemplates, output.QuizTemplateOutput{
			QuestionId:     question.Id,
			SelectedOption: answer.Option,
			CorrectOption:  correctAlternative.Option,
		})
	}

	userScore := uc.calculateUserScore(rightsCount, len(input.Answers))
	uc.saveQuizScore(input.User, userScore)
	rate := uc.calculateSuccessRate(userScore)

	return &output.QuizOutput{
		Resume:       fmt.Sprintf("You were better than %.f%% of all quizzers", rate),
		RightAnswers: rightsCount,
		WrongAnswers: wrongsCount,
		QuizTemplate: quizTemplates,
	}, nil
}

func (uc *correctQuizUseCase) calculateUserScore(rightsCount, totalAnswers int) float64 {
	return (float64(rightsCount) / float64(totalAnswers)) * 10
}

func (uc *correctQuizUseCase) calculateSuccessRate(userScore float64) float64 {
	scores := uc.quizRepository.GetAllScores()
	sort.Float64s(*scores)

	position := uc.binarySearch(*scores, userScore)
	return (float64(position) / float64(len(*scores))) * 100.0
}

func (uc *correctQuizUseCase) saveQuizScore(userName string, score float64) {
	uc.quizRepository.Save(entities.QuizScore{
		UserName: userName,
		Score:    score,
	})
}

// binarySearch returns the index of the userScore in the scores array
func (uc *correctQuizUseCase) binarySearch(scores []float64, userScore float64) int {
	low := 0
	high := len(scores) - 1
	for low <= high {
		mid := low + (high-low)/2
		if scores[mid] == userScore {
			return mid
		}
		if scores[mid] < userScore {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return low
}
