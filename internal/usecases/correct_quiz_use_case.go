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
func (g *correctQuizUseCase) Execute(input input.QuizInput) (*output.QuizOutput, error) {
	questionIDs := make(map[int]bool)
	rightsCount := 0
	wrongsCount := 0
	var quizTemplates []output.QuizTemplateOutput

	for _, answer := range input.Answers {
		if _, exists := questionIDs[answer.QuestionId]; exists {
			return nil, fmt.Errorf("duplicate question ID: %d", answer.QuestionId)
		}

		questionIDs[answer.QuestionId] = true

		question, error := g.questionsRepository.FindQuestionById(answer.QuestionId)

		if error != nil {
			return nil, error
		}

		correctAlternative, error := question.GetCorrectAlternative()

		if error != nil {
			return nil, error
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

	userScore := (float64(rightsCount) / float64(len(input.Answers))) * 10

	scores := g.quizRepository.GetAllScores()
	sort.Float64s(*scores)

	position := g.binarySearch(*scores, userScore)

	rate := (float64(position) / float64(len(*scores)) * 100.0)

	g.quizRepository.Save(entities.QuizScore{
		UserName: input.User,
		Score:    userScore,
	})

	return &output.QuizOutput{
		Resume: fmt.Sprintf(
			"You answered %d question correctly out of %d. "+
				"You made %d error. "+
				"You were better than %.f%% of all quizzers",
			rightsCount,
			len(input.Answers),
			wrongsCount,
			rate,
		),
		RightAnswers: rightsCount,
		WrongAnswers: wrongsCount,
		QuizTemplate: quizTemplates,
	}, nil
}

// binarySearch returns the index of the userScore in the scores array
func (g *correctQuizUseCase) binarySearch(scores []float64, userScore float64) int {
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
