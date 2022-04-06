package services

import (
	"fmt"
	"testing"

	"github.com/doffy007/golang-unit-test/entities"
	"github.com/doffy007/golang-unit-test/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var categoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryService = CategoryService{Repository: categoryRepository}

func TestCategoryService_GetNotFound(t *testing.T) {
	categoryRepository.Mock.On("FindById", 1).Return(nil)

	category, err := categoryService.Get(1)
	assert.Nil(t, category)
	assert.NotNil(t, err)

	fmt.Println(err)
}

func TestCategoryService_GetSuccess(t *testing.T) {
	category := entities.Category{
		Id:   4,
		Name: "Laptop",
	}
	categoryRepository.Mock.On("FindById", 10).Return(category) 
	// apapun nilai yang ada pada id category jika selama dalam bentuk int = aman
	// jika kategori repositori mock tidak sama dengan result maka akan error
	result, err := categoryService.Get(10)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, category.Id, result.Id)
	assert.Equal(t, category.Name, result.Name)

	fmt.Println(category, err)
}