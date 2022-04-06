package repository

import "github.com/doffy007/golang-unit-test/entities"

type CategoryRepository interface {
	FindById(id int) *entities.Category
}