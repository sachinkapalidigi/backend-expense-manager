package categoriescontroller

import (
	"net/http"
	"strconv"

	"github.com/sachinkapalidigi/backend-expense-manager/services"

	"github.com/sachinkapalidigi/backend-expense-manager/domain/categories"
	"github.com/sachinkapalidigi/backend-expense-manager/utils/errors"

	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	var category categories.Category
	if err := c.BindJSON(&category); err != nil {
		restErr := errors.NewBadRequestError("Invalid JSON body")
		c.JSON(restErr.Status, restErr)
		return
	}
	result, createErr := services.CategoriesService.CreateCategory(category)
	if createErr != nil {
		c.JSON(createErr.Status, createErr)
		return
	}

	c.JSON(http.StatusOK, result)
	return
}

func Get(c *gin.Context) {
	categoryId, err := strconv.ParseInt(c.Param("category_id"), 10, 64)
	if err != nil {
		restErr := errors.NewBadRequestError("Category Id must be an integer")
		c.JSON(restErr.Status, restErr)
	}

	result, getErr := services.CategoriesService.GetCategory(categoryId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	c.JSON(http.StatusOK, result)
}

func GetAll(c *gin.Context) {

	c.JSON(http.StatusOK, map[string]string{"message": "Not implemented"})
}
