// Code generated by mockery v2.42.3. DO NOT EDIT.

package mocks

import (
	entities "quiz-fast-track/internal/entities"

	mock "github.com/stretchr/testify/mock"
)

// QuizRepository is an autogenerated mock type for the QuizRepository type
type QuizRepository struct {
	mock.Mock
}

// GetAllScores provides a mock function with given fields:
func (_m *QuizRepository) GetAllScores() *[]float64 {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetAllScores")
	}

	var r0 *[]float64
	if rf, ok := ret.Get(0).(func() *[]float64); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]float64)
		}
	}

	return r0
}

// Save provides a mock function with given fields: quizScore
func (_m *QuizRepository) Save(quizScore entities.QuizScore) error {
	ret := _m.Called(quizScore)

	if len(ret) == 0 {
		panic("no return value specified for Save")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(entities.QuizScore) error); ok {
		r0 = rf(quizScore)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewQuizRepository creates a new instance of QuizRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewQuizRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *QuizRepository {
	mock := &QuizRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
