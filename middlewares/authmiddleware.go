package middlewares

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sachinkapalidigi/backend-expense-manager/domain/users"
	"github.com/sachinkapalidigi/backend-expense-manager/utils/accesstokenutil"
	"github.com/sachinkapalidigi/backend-expense-manager/utils/errors"
)

var (
	AuthMiddleware authMiddlewareInterface = &authMiddleware{}
)

type authMiddleware struct{}

type authMiddlewareInterface interface {
	UserLoader() gin.HandlerFunc
	EnforceAuthenticatedMiddleware() gin.HandlerFunc
}

func (a *authMiddleware) UserLoader() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearer := c.Request.Header.Get("Authorization")
		fmt.Println(bearer)
		if bearer != "" {
			// jwt decode

			user := accesstokenutil.ParseToken(bearer)
			fmt.Println(user)
			// get user data from user id
			if user != nil {
				c.Set("currentUser", user)
				c.Set("currentUserId", user.ID)
			}
		}
	}
}

// EnforceAuthenticatedMiddleware : Reject if unauthorized
func (a *authMiddleware) EnforceAuthenticatedMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, exists := c.Get("currentUser")
		fmt.Println(user)
		if exists && user.(*users.User).ID != 0 {
			return
		} else {
			err, exists := c.Get("authErr")
			if exists {
				c.AbortWithStatusJSON(http.StatusForbidden, err)
			} else {
				restErr := errors.NewNotAuthorizedError("Invalid or No token present for authentication")
				c.JSON(restErr.Status, restErr)
				c.Abort()
			}
		}
	}
}
