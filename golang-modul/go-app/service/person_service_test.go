package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/ugik-dev/go-learn-app/models"
	"github.com/ugik-dev/go-learn-app/repository"
)

var personRepo = &repository.PersonRepoMock{Mock: mock.Mock{}}
var personService = PersonService{Repository: personRepo}

func TestPersonService_GetNotFound(t *testing.T) {

	// program mock

	personRepo.Mock.On("FindById", "1").Return(nil)

	result, err := personService.Get("1")
	assert.Nil(t, result)
	assert.NotNil(t, err)

}

func TestPersonService_GetSuccess(t *testing.T) {

	// program mock

	person := models.Person{
		Id:      "1",
		Name:    "Zea",
		Age:     3,
		Address: "pringsewu",
	}

	personRepo.Mock.On("FindById", "2").Return(person)
	result, err := personService.Get("2")
	assert.Nil(t, err)
	assert.NotNil(t, result)

}
