package categoriescontroller

import (
	"net/http"
	"strconv"

	"github.com/sachinkapalidigi/backend-expense-manager/services"

	"github.com/sachinkapalidigi/backend-expense-manager/domain/categories"
	"github.com/sachinkapalidigi/backend-expense-manager/domain/users"
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
	user, exists := c.Get("currentUser")
	if !exists {
		restErr := errors.NewNotAuthorizedError("cannot create categories, Login to create!")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, createErr := services.CategoriesService.CreateCategory(category, user.(*users.User).ID)
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
	user, exists := c.Get("currentUser")
	if !exists {
		restErr := errors.NewNotAuthorizedError("cannot get category, Login to create!")
		c.JSON(restErr.Status, restErr)
		return
	}
	result, getErr := services.CategoriesService.GetCategory(categoryId, user.(*users.User).ID)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	c.JSON(http.StatusOK, result)
}

// GetAll : Get all the categories
func GetAll(c *gin.Context) {
	user, exists := c.Get("currentUser")
	if !exists {
		restErr := errors.NewNotAuthorizedError("cannot get categories, Login to create!")
		c.JSON(restErr.Status, restErr)
		return
	}
	results, err := services.CategoriesService.GetAllCategories(user.(*users.User).ID)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	ct := map[string]categories.Categories{"categories": results}
	c.JSON(http.StatusOK, ct)
}
