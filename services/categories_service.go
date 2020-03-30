package services

import (
	"github.com/sachinkapalidigi/backend-expense-manager/domain/categories"
	"github.com/sachinkapalidigi/backend-expense-manager/utils/dateutil"
	"github.com/sachinkapalidigi/backend-expense-manager/utils/errors"
)

var (
	// CategoriesService : interface for methods in categories service
	CategoriesService categoriesServiceInterface = &categoriesService{}
)

type categoriesService struct{}

type categoriesServiceInterface interface {
	CreateCategory(categories.Category) (*categories.Category, *errors.RestErr)
	GetCategory(int64) (*categories.Category, *errors.RestErr)
	GetAllCategories() (categories.Categories, *errors.RestErr)
}

func (s *categoriesService) CreateCategory(category categories.Category) (*categories.Category, *errors.RestErr) {
	// validate
	if err := category.Validate(); err != nil {
		return nil, err
	}
	category.CreatedAt = dateutil.GetNowDBFormat()
	// store in table
	if err := category.Create(); err != nil {
		return nil, err
	}
	// return created category
	return &category, nil
}

func (s *categoriesService) GetCategory(categoryId int64) (*categories.Category, *errors.RestErr) {

	result := categories.Category{ID: categoryId}

	if err := result.Get(); err != nil {
		return nil, err
	}

	return &result, nil
}

func (s *categoriesService) GetAllCategories() (categories.Categories, *errors.RestErr) {
	c := categories.Category{}
	return c.GetAll()
}
