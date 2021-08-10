package repositories

import "golang-unit-test/entities"

type CategoryRepository interface {
	FindById(id string) *entities.Category
}
