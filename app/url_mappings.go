package app

import (
	"net/http"

	"github.com/sachinkapalidigi/backend-expense-manager/controllers/expensescontroller"

	"github.com/sachinkapalidigi/backend-expense-manager/controllers/categoriescontroller"

	"github.com/gin-gonic/gin"
)

func mapUrls() {

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]string{
			"message": "Server is working fine",
		})
	})

	router.POST("/categories", categoriescontroller.Create)
	router.GET("/categories/:category_id", categoriescontroller.Get)
	router.GET("/categories", categoriescontroller.GetAll)

	router.POST("/expenses", expensescontroller.Create)
	router.GET("/expenses/:expense_id", expensescontroller.Get)
	router.GET("/expenses", expensescontroller.GetAll)
}
