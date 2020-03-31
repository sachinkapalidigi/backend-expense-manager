package expenses

import (
	"strings"

	"github.com/sachinkapalidigi/backend-expense-manager/domain/categories"
	"github.com/sachinkapalidigi/backend-expense-manager/utils/errors"
)

// Expense : structure for expense
type Expense struct {
	ID          int64  `json:"id"`
	CategoryID  int64  `json:"category_id"`
	Amount      string `json:"amount"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	PaymentMode string `json:"payment_mode"`
}

// ExpenseCategory : structure for expense with category
type ExpenseCategory struct {
	Expense  Expense
	Category categories.Category
}

// Expenses : structure for expenses
type Expenses []ExpenseCategory

// Validate : validate the expense
func (e *Expense) Validate() *errors.RestErr {
	e.Amount = strings.TrimSpace(e.Amount)
	if e.Amount == "" {
		return errors.NewBadRequestError("Invalid Amount")
	}
	if e.CategoryID == 0 {
		return errors.NewBadRequestError("Invalid Category")
	}
	e.PaymentMode = strings.TrimSpace(e.PaymentMode)
	if e.PaymentMode == "" {
		return errors.NewBadRequestError("Invalid Payment mode")
	}
	return nil
}
