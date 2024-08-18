package repository

import (
	"github.com/stretchr/testify/mock"
	"github.com/ugik-dev/go-learn-app/models"
)

type PersonRepoMock struct {
	Mock mock.Mock
}

func (repository *PersonRepoMock) FindById(id string) *models.Person {
	arguments := repository.Mock.Called(id)
	if arguments.Get(0) == nil {
		return nil
	} else {
		person := arguments.Get(0).(models.Person)
		return &person
	}
}
