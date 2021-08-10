package repositories

import (
	"golang-unit-test/entities"

	"github.com/stretchr/testify/mock"
)

type CategoryRepositoryMock struct {
	Mock mock.Mock
}

func (crm *CategoryRepositoryMock) FindById(id string) *entities.Category {
	arguments := crm.Mock.Called(id)

	if arguments.Get(0) == nil {
		return nil
	} else {
		category := arguments.Get(0).(entities.Category)
		return &category
	}
}
