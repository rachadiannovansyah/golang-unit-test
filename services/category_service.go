package services

import (
	"errors"
	"golang-unit-test/entities"
	"golang-unit-test/repositories"
)

type CategoryService struct {
	Repository repositories.CategoryRepository
}

func (cs CategoryService) Get(id string) (*entities.Category, error) {
	category := cs.Repository.FindById(id)
	if category == nil {
		return nil, errors.New("Category not found")
	} else {
		return category, nil
	}
}
