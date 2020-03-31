package expenses

import (
	"strings"

	"github.com/sachinkapalidigi/backend-expense-manager/domain/categories"

	"github.com/sachinkapalidigi/backend-expense-manager/datasources/postgresql/expensesdb"
	"github.com/sachinkapalidigi/backend-expense-manager/logger"
	"github.com/sachinkapalidigi/backend-expense-manager/utils/errors"
)

const (
	queryInsertExpense                          = "INSERT INTO expenses(category_id, amount, description, created_at, payment_mode) VALUES($1, $2, $3, $4, $5) RETURNING id;"
	querySelectExpenseById                      = "SELECT expenses.id, expenses.category_id, expenses.amount, expenses.description, expenses.created_at, expenses.payment_mode, categories.id, categories.category_name, categories.description, categories.created_at FROM expenses INNER JOIN categories ON expenses.category_id = categories.id WHERE expenses.id=$1;"
	querySelectExpensesByCategoryIdBetweenDates = "SELECT expenses.id, expenses.category_id, expenses.amount, expenses.description, expenses.created_at, expenses.payment_mode, categories.id, categories.category_name, categories.description, categories.created_at FROM expenses INNER JOIN categories ON expenses.category_id = categories.id WHERE categories.id=$1 AND expenses.created_at BETWEEN $2 AND $3 ORDER BY expenses.created_at DESC;"
	querySelectExpensesBetweenDates             = "SELECT expenses.id, expenses.category_id, expenses.amount, expenses.description, expenses.created_at, expenses.payment_mode, categories.id, categories.category_name, categories.description, categories.created_at FROM expenses INNER JOIN categories ON expenses.category_id = categories.id WHERE expenses.created_at BETWEEN $1 AND $2 ORDER BY expenses.created_at DESC;"
)

// AddExpense : Save to database
func (expense *Expense) AddExpense() *errors.RestErr {
	stmt, err := expensesdb.Client.Prepare(queryInsertExpense)
	if err != nil {
		logger.Error("Error in statement", err)
		return errors.NewInternalServerError("Couldnot save expense")
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

// GetExpenseDetails : get the expensedetails from expense id
func (expense *Expense) GetExpenseDetails() (*ExpenseCategory, *errors.RestErr) {
	stmt, err := expensesdb.Client.Prepare(querySelectExpenseById)
	if err != nil {
		logger.Error("Error in statement", err)
		return nil, errors.NewInternalServerError("Couldnot get expense")
	}

	defer stmt.Close()

	var category categories.Category
	result := stmt.QueryRow(&expense.ID)
	if err := result.Scan(&expense.ID, &expense.CategoryID, &expense.Amount, &expense.Description, &expense.CreatedAt, &expense.PaymentMode, &category.ID, &category.CategoryName, &category.Description, &category.CreatedAt); err != nil {
		if strings.Contains(err.Error(), "no rows") {
			return nil, errors.NewNotFoundError("expenses not found")
		}
		return nil, errors.NewInternalServerError("Couldnot retrieve expenses")
	}
	var expenseCategory = ExpenseCategory{
		Expense:  *expense,
		Category: category,
	}
	return &expenseCategory, nil
}

func (expense *Expense) GetAll() ([]ExpenseCategory, *errors.RestErr) {

	return nil, nil
}

// GetAllByDetails : based on category id and dates in between gives slice of Expenses
func (e *Expense) GetAllByDetails(categoryId int64, from string, to string) ([]ExpenseCategory, *errors.RestErr) {

	if categoryId != 0 {
		stmt, err := expensesdb.Client.Prepare(querySelectExpensesByCategoryIdBetweenDates)
		if err != nil {
			logger.Error("Error in statement", err)
			return nil, errors.NewInternalServerError("Couldnot get expense")
		}
		defer stmt.Close()
		rows, err := stmt.Query(categoryId, from, to)
		if err != nil {
			logger.Error("Error in statement", err)
			return nil, errors.NewInternalServerError("Couldnot get expense")
		}
		defer rows.Close()
		var expenses []ExpenseCategory

		for rows.Next() {
			var expense Expense
			var category categories.Category
			if err := rows.Scan(&expense.ID, &expense.CategoryID, &expense.Amount, &expense.Description, &expense.CreatedAt, &expense.PaymentMode, &category.ID, &category.CategoryName, &category.Description, &category.CreatedAt); err != nil {
				return nil, errors.NewInternalServerError("Parse Error")
			}
			var expenseCategory = ExpenseCategory{
				Expense:  expense,
				Category: category,
			}
			expenses = append(expenses, expenseCategory)
		}
		if len(expenses) == 0 {
			return nil, errors.NewNotFoundError("No expenses for the given category")
		}

		return expenses, nil
	} else {
		stmt, err := expensesdb.Client.Prepare(querySelectExpensesBetweenDates)
		if err != nil {
			logger.Error("Error in statement", err)
			return nil, errors.NewInternalServerError("Couldnot get expense")
		}
		defer stmt.Close()
		rows, err := stmt.Query(from, to)
		if err != nil {
			logger.Error("Error in statement", err)
			return nil, errors.NewInternalServerError("Couldnot get expense")
		}
		defer rows.Close()
		var expenses []ExpenseCategory

		for rows.Next() {
			var expense Expense
			var category categories.Category
			if err := rows.Scan(&expense.ID, &expense.CategoryID, &expense.Amount, &expense.Description, &expense.CreatedAt, &expense.PaymentMode, &category.ID, &category.CategoryName, &category.Description, &category.CreatedAt); err != nil {
				return nil, errors.NewInternalServerError("Parse Error")
			}
			var expenseCategory = ExpenseCategory{
				Expense:  expense,
				Category: category,
			}
			expenses = append(expenses, expenseCategory)
		}
		if len(expenses) == 0 {
			return nil, errors.NewNotFoundError("No expenses for the given category")
		}
		return expenses, nil
	}

}
