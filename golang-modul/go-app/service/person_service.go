package service

import (
	"errors"

	"github.com/ugik-dev/go-learn-app/models"
	"github.com/ugik-dev/go-learn-app/repository"
)

type PersonService struct {
	Repository repository.PersonRepo
}

func (service PersonService) Get(id string) (*models.Person, error) {
	person := service.Repository.FindById(id)
	if person == nil {
		return nil, errors.New("PERSON NOT FOUND")
	} else {
		return person, nil
	}
}
