package services

import (
	"errors"

	"github.com/doffy007/golang-unit-test/entities"
	"github.com/doffy007/golang-unit-test/repository"
)

type CategoryService struct {
	Repository repository.CategoryRepository
}

func (service CategoryService) Get(id int) (*entities.Category, error) {
	category := service.Repository.FindById(id)
	if category == nil {
		return nil, errors.New("category Not Found")
	} else {
		return category, nil
	}
}