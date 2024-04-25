package usecases

import (
	"quiz-fast-track/internal/usecases/contracts"
	"quiz-fast-track/internal/usecases/ports/output"
	"quiz-fast-track/internal/usecases/repositories"
)

type getQuestionsUseCase struct {
	questionsRepository repositories.QuestionsRepository
}

func NewGetQuestionsUseCase(questionsRepository repositories.QuestionsRepository) contracts.GetQuestionsUseCase {
	return &getQuestionsUseCase{
		questionsRepository: questionsRepository,
	}
}

// Composition estou atrelando a função ao objeto
func (g *getQuestionsUseCase) Execute() []output.QuestionOutput {
	questions := g.questionsRepository.GetAll()

	var questionsOutput []output.QuestionOutput

	for _, question := range *questions {
		var alternativesOutPut []output.AlternativeOutput
		for _, alternative := range question.Alternatives {
			alternativeOutPut := output.AlternativeOutput{
				Id:     alternative.Id,
				Option: alternative.Option,
				Text:   alternative.Text,
			}
			alternativesOutPut = append(alternativesOutPut, alternativeOutPut)
		}
		questionOutput := output.QuestionOutput{
			Id:           question.Id,
			Text:         question.Text,
			Alternatives: alternativesOutPut,
		}
		questionsOutput = append(questionsOutput, questionOutput)
	}
	return questionsOutput
}
