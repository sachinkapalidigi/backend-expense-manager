package userscontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sachinkapalidigi/backend-expense-manager/domain/users"
	"github.com/sachinkapalidigi/backend-expense-manager/services"
	"github.com/sachinkapalidigi/backend-expense-manager/utils/errors"
)

// RegisterUser : Adds new User
func RegisterUser(c *gin.Context) {
	var user users.User
	var err error
	if err = c.BindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("Invalid JSON body")
		c.JSON(restErr.Status, restErr)
		return
	}

	_, restErr := services.UsersServices.Create(user)
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}

	c.JSON(http.StatusOK, map[string]string{
		"message": "Successfully Registered User",
	})
	return
}

// LoginUser : Login user based on email and password
func LoginUser(c *gin.Context) {
	// get email and password
	var user *users.User
	if err := c.BindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("Invalid JSON body")
		c.JSON(restErr.Status, restErr)
		return
	}
	// service to verify
	accessToken, restErr := services.UsersServices.Authenticate(user)
	// generate token and send back
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}

	c.JSON(http.StatusOK, accessToken)
	return
}
