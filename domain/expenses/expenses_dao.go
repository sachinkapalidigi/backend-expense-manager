package expenses

import (
	"strings"

	"github.com/sachinkapalidigi/backend-expense-manager/datasources/postgresql/expensesdb"
	"github.com/sachinkapalidigi/backend-expense-manager/logger"
	"github.com/sachinkapalidigi/backend-expense-manager/utils/errors"
)

const (
	queryInsertExpense = "INSERT INTO expenses(category_id, amount, description, created_at, payment_mode) VALUES($1, $2, $3, $4, $5) RETURNING id;"
)

// AddExpense : Save to database
func (expense *Expense) AddExpense() *errors.RestErr {
	stmt, err := expensesdb.Client.Prepare(queryInsertExpense)
	if err != nil {
		logger.Error("Error in statement", err)
		return errors.NewInternalServerError("Couldnot save expense to database")
	}
	defer stmt.Close()
	if err := stmt.QueryRow(&expense.CategoryID, &expense.Amount, &expense.Description, &expense.CreatedAt, &expense.PaymentMode).Scan(&expense.ID); err != nil {
		logger.Error("expense couldnot be created", err)
		if strings.Contains(err.Error(), "violates foreign key constraint") {
			return errors.NewBadRequestError("Invalid Category")
		}
		return errors.NewInternalServerError("Could not add expense")
	}
	return nil
}
