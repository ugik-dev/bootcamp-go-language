package repository

import "github.com/ugik-dev/go-learn-app/models"

type PersonRepo interface {
	FindById(id string) *models.Person
}
