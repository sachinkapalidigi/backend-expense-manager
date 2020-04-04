package services

import (
	"github.com/sachinkapalidigi/backend-expense-manager/domain/expenses"
	"github.com/sachinkapalidigi/backend-expense-manager/utils/dateutil"
	"github.com/sachinkapalidigi/backend-expense-manager/utils/errors"
)

var (
	ExpensesService expensesServiceInterface = &expensesService{}
)

type expensesServiceInterface interface {
	AddExpense(expenses.Expense, int64) (*expenses.Expense, *errors.RestErr)
	GetExpense(int64, int64) (*expenses.ExpenseCategory, *errors.RestErr)
	GetExpenses(int64, string, string, int64) ([]expenses.ExpenseCategory, *errors.RestErr)
}

type expensesService struct{}

func (s *expensesService) AddExpense(expense expenses.Expense, userID int64) (*expenses.Expense, *errors.RestErr) {
	if err := expense.Validate(); err != nil {
		return nil, err
	}
	expense.CreatedAt = dateutil.GetNowDBFormat()
	if err := expense.AddExpense(userID); err != nil {
		return nil, err
	}

	return &expense, nil
}

func (s *expensesService) GetExpense(expenseID int64, userID int64) (*expenses.ExpenseCategory, *errors.RestErr) {
	var expense = expenses.Expense{ID: expenseID}
	result, err := expense.GetExpenseDetails(userID)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *expensesService) GetExpenses(categoryId int64, from string, to string, userID int64) ([]expenses.ExpenseCategory, *errors.RestErr) {
	var expense expenses.Expense
	return expense.GetAllByDetails(categoryId, from, to, userID)
}
