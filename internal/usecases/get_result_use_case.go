package usecases

import (
	"fmt"
	"quiz-fast-track/internal/usecases/contracts"
	"quiz-fast-track/internal/usecases/ports/input"
	"quiz-fast-track/internal/usecases/ports/output"
	"quiz-fast-track/internal/usecases/repositories"
)


type getResultUseCase struct {
	questionsRepository repositories.QuestionsRepository
}

func GetResultUseCase(questionsRepository repositories.QuestionsRepository) contracts.GetResultUseCase {
	return &getResultUseCase{
		questionsRepository: questionsRepository,
	}
}

func (g *getResultUseCase) Execute(answers input.AnswersInput) (output.ResultOutput, error) {
	questionIDs := make(map[int]bool)
	correctAnswersCount := 0
	errorCount := 0
	var quizTemplates []output.QuizTemplateOutput


	for _, answer := range answers.Answers {
		if _, exists := questionIDs[answer.QuestionId]; exists {
			return output.ResultOutput{}, fmt.Errorf("duplicate question ID: %d", answer.QuestionId)
		}
		questionIDs[answer.QuestionId] = true
		correctAnswer, err := g.questionsRepository.GetAnswerById(answer.QuestionId)
		if err != nil {
			return output.ResultOutput{}, fmt.Errorf("error retrieving answer for question ID %d: %v", answer.QuestionId, err)
		}
		
		// User should be able to answer all the questions and then post his/hers answers and get back how many correct answers they had, displayed to the user.
		// na minha funçao do repository eu vou mandar a resposta do usuario e vou receber a alternativa corret
		if answer.Option == correctAnswer.Option {
			correctAnswersCount++
		} else {
			errorCount++
		}

		quizTemplates = append(quizTemplates, output.QuizTemplateOutput{
			QuestionId:     answer.QuestionId,
			SelectedOption: answer.Option,
			CorrectOption:  correctAnswer.Option,
		})
	}

	resume := fmt.Sprintf("You answered %d questions correctly out of %d. You made %d errors.", correctAnswersCount, len(answers.Answers), errorCount)

	result := output.ResultOutput{
		Resume:      resume,
		Hits:        correctAnswersCount,
		Erros:       errorCount,
		QuizTemplate: quizTemplates,
	}
	return result, nil
}

