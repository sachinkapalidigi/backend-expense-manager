package expensescontroller

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/sachinkapalidigi/backend-expense-manager/utils/dateutil"

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
	var expenseID int64
	var parseErr error
	expenseID, parseErr = strconv.ParseInt(c.Param("expense_id"), 10, 64)
	if parseErr != nil {
		restErr := errors.NewBadRequestError("Invalid expense id")
		c.JSON(restErr.Status, restErr)
		return
	}
	result, err := services.ExpensesService.GetExpense(expenseID)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, result)
	return
}

// GetAll : get all expenses controller based on query params
func GetAll(c *gin.Context) {
	var categoryID int64
	var err error
	if c.Query("category_id") == "all" || c.Query("category_id") == "" {
		categoryID = 0
	} else {
		categoryID, err = strconv.ParseInt(c.Query("category_id"), 10, 64)
		if err != nil {
			restErr := errors.NewBadRequestError("Invalid category Id")
			c.JSON(restErr.Status, restErr)
			return
		}
	}
	from := strings.TrimSpace(c.Query("from"))
	to := strings.TrimSpace(c.Query("to"))
	if from == "" {
		from = "2020-03-30"
	}
	if to == "" {
		to = dateutil.GetNowDBFormat()
	}

	result, restErr := services.ExpensesService.GetExpenses(categoryID, from, to)
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}
	c.JSON(http.StatusOK, result)
	return
}
