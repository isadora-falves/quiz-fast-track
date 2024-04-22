package usecases

import (
	"fmt"
	"quiz-fast-track/internal/usecases/contracts"
	"quiz-fast-track/internal/usecases/ports/input"
	"quiz-fast-track/internal/usecases/ports/output"
	"quiz-fast-track/internal/usecases/repositories"
)

type correctQuizUseCase struct {
	questionsRepository repositories.QuestionsRepository
}

func NewCorrectQuizUseCase(questionsRepository repositories.QuestionsRepository) contracts.CorrectQuizUseCase {
	return &correctQuizUseCase{
		questionsRepository: questionsRepository,
	}
}

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

		question, _ := g.questionsRepository.FindQuestionById(answer.QuestionId)

		correctAlternative, _ := question.GetCorrectAlternative()


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


		// if err != nil {
		// 	return output.ResultOutput{}, fmt.Errorf("error retrieving answer for question ID %d: %v", answer.QuestionId, err)
		// }

		// // User should be able to answer all the questions and then post his/hers answers and get back how many correct answers they had, displayed to the user.
		// if answer.Option == correctAnswer.Option {
		// 	correctAnswersCount++
		// } else {
		// 	errorCount++
		// }

		// quizTemplates = append(quizTemplates, output.QuizTemplateOutput{
		// 	QuestionId:     answer.QuestionId,
		// 	SelectedOption: answer.Option,
		// 	CorrectOption:  correctAnswer.Option,
		// })
	}

	// resume := fmt.Sprintf("You answered %d questions correctly out of %d. You made %d errors.", correctAnswersCount, len(answers.Answers), errorCount)

	// result := output.ResultOutput{
	// 	Resume:       resume,
	// 	Hits:         correctAnswersCount,
	// 	Erros:        errorCount,
	// 	QuizTemplate: quizTemplates,
	// }
	return &output.QuizOutput{
		Resume:       fmt.Sprintf("You answered %d question correctly out of %d. You made %d error.", rightsCount, len(input.Answers), wrongsCount),
		RightAnswers: rightsCount,
		WrongAnswers: wrongsCount,
		QuizTemplate: quizTemplates,
	}, nil
}
