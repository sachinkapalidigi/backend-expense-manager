package categories

import (
	"fmt"

	"github.com/sachinkapalidigi/backend-expense-manager/datasources/postgresql/expensesdb"
	"github.com/sachinkapalidigi/backend-expense-manager/logger"
	"github.com/sachinkapalidigi/backend-expense-manager/utils/errors"
)

const (
	queryCreateCategory = "INSERT INTO categories(category_name, description, created_at) VALUES($1, $2, $3) RETURNING id;"
)

// Create : Add category to database
func (category *Category) Create() *errors.RestErr {
	fmt.Println(expensesdb.Client)
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
