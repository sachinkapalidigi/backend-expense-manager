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
	AddExpense(expenses.Expense) (*expenses.Expense, *errors.RestErr)
	GetExpense(int64) (*expenses.ExpenseCategory, *errors.RestErr)
	GetExpenses(int64, string, string) ([]expenses.ExpenseCategory, *errors.RestErr)
}

type expensesService struct{}

func (s *expensesService) AddExpense(expense expenses.Expense) (*expenses.Expense, *errors.RestErr) {
	if err := expense.Validate(); err != nil {
		return nil, err
	}
	expense.CreatedAt = dateutil.GetNowDBFormat()
	if err := expense.AddExpense(); err != nil {
		return nil, err
	}

	return &expense, nil
}

func (s *expensesService) GetExpense(expenseID int64) (*expenses.ExpenseCategory, *errors.RestErr) {
	var expense = expenses.Expense{ID: expenseID}
	result, err := expense.GetExpenseDetails()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *expensesService) GetExpenses(categoryId int64, from string, to string) ([]expenses.ExpenseCategory, *errors.RestErr) {
	var expense expenses.Expense
	return expense.GetAllByDetails(categoryId, from, to)
}
