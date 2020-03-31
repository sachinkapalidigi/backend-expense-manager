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
