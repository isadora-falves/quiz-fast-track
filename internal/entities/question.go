package entities

import "fmt"
type Question struct {
	Id           int
	Text         string
	Alternatives []Alternative
}

func (q *Question) GetCorrectAlternative() (*Alternative, error) {
	for _, alternative := range q.Alternatives {
		if alternative.IsCorrect {
			return &alternative, nil
		}
	}

	return nil, fmt.Errorf("there is no alternative correct for this question")
}
