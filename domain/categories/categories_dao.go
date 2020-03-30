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
	querySelectCategoryById = "SELECT id, category_name, description, created_at FROM categories WHERE id=$1;"
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
		restErr := errors.NewInternalServerError("Category could not be saved")
		return restErr
	}

	return nil
}

// Get : Get category based on user id
func (category *Category) Get() *errors.RestErr {
	stmt, err := expensesdb.Client.Prepare(querySelectCategoryById)

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
