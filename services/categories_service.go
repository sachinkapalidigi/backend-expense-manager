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
