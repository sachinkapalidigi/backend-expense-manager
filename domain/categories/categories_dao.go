package categories

import (
	"fmt"
	"strings"

	"github.com/sachinkapalidigi/backend-expense-manager/datasources/postgresql/expensesdb"
	"github.com/sachinkapalidigi/backend-expense-manager/logger"
	"github.com/sachinkapalidigi/backend-expense-manager/utils/errors"
)

const (
	queryCreateCategory     = "INSERT INTO categories(category_name, description, created_at) VALUES($1, $2, $3) RETURNING id;"
	querySelectCategoryByID = "SELECT id, category_name, description, created_at FROM categories WHERE id=$1;"
	querySelectCategories   = "SELECT id, category_name, description, created_at FROM categories"
)

// Create : Add category to database
func (category *Category) Create() *errors.RestErr {

	// prepare statement
	stmt, err := expensesdb.Client.Prepare(queryCreateCategory)
	if err != nil {
		fmt.Println(err)
		logger.Error("Error in preparing statement", err)
		restErr := errors.NewInternalServerError("Error in connecting to database")
		return restErr
	}
	// defer : close statement
	defer stmt.Close()
	// execute statement
	err = stmt.QueryRow(&category.CategoryName, &category.Description, &category.CreatedAt).Scan(&category.ID)
	if err != nil {
		logger.Error("Error in saving to database", err)
		if strings.Contains(err.Error(), "duplicate key value") {
			return errors.NewBadRequestError("Category already exists")
		}
		restErr := errors.NewInternalServerError("Category could not be saved")
		return restErr
	}

	return nil
}

// Get : Get category based on user id
func (category *Category) Get() *errors.RestErr {
	stmt, err := expensesdb.Client.Prepare(querySelectCategoryByID)

	if err != nil {
		logger.Error("Error in preparing statement", err)
		return errors.NewInternalServerError("Error while getting category")
	}
	defer stmt.Close()

	result := stmt.QueryRow(&category.ID)
	if err := result.Scan(&category.ID, &category.CategoryName, &category.Description, &category.CreatedAt); err != nil {
		if strings.Contains(err.Error(), "no rows") {
			return errors.NewNotFoundError("category not found")
		}
		return errors.NewInternalServerError("Couldnot retrieve category")
	}

	return nil
}

// GetAll : Get all the categories
func (category *Category) GetAll() (Categories, *errors.RestErr) {
	stmt, err := expensesdb.Client.Prepare(querySelectCategories)
	if err != nil {
		logger.Error("problem with statement", err)
		return nil, errors.NewInternalServerError("Could not fetch categories")
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, errors.NewInternalServerError("Could not fetch categories")
	}
	defer rows.Close()

	results := make(Categories, 0)
	for rows.Next() {
		var category Category
		if err := rows.Scan(&category.ID, &category.CategoryName, &category.Description, &category.CreatedAt); err != nil {
			return nil, errors.NewInternalServerError("parse error")
		}
		results = append(results, category)
	}

	if len(results) == 0 {
		return nil, errors.NewNotFoundError("No categories added")
	}

	return results, nil
}
