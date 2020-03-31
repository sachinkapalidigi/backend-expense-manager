package expensescontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sachinkapalidigi/backend-expense-manager/domain/expenses"
	"github.com/sachinkapalidigi/backend-expense-manager/services"
	"github.com/sachinkapalidigi/backend-expense-manager/utils/errors"
)

// Create : add expense controller
func Create(c *gin.Context) {
	var expense expenses.Expense
	if err := c.BindJSON(&expense); err != nil {
		restErr := errors.NewBadRequestError("Invalid JSON")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, restErr := services.ExpensesService.AddExpense(expense)
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}
	c.JSON(http.StatusOK, result)
	return
}

// Get : get expense from expense id controller
func Get(c *gin.Context) {

}

// GetAll : get all expenses controller based on query params
func GetAll(c *gin.Context) {

}
