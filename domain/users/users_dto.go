package users

import (
	"strings"

	"github.com/sachinkapalidigi/backend-expense-manager/utils/errors"
	"gopkg.in/go-playground/validator.v9"
)

// User : type user structure
type User struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,passwd"`
	CreatedAt string `json:"created_at"`
}

// register all custom rules and keep the validatory in a util

// Validate : User info validate
func (user *User) Validate() *errors.RestErr {
	user.Name = strings.TrimSpace(user.Name)
	if user.Name == "" {
		return errors.NewBadRequestError("Name is required")
	}
	v := validator.New()
	_ = v.RegisterValidation("passwd", func(fl validator.FieldLevel) bool {
		return len(fl.Field().String()) > 6
	})

	err := v.Struct(user)

	if err != nil {
		return errors.NewBadRequestError(err.Error())
	}

	return nil
}

// ValidateLogin : Validation for login request
func (user *User) ValidateLogin() *errors.RestErr {
	v := validator.New()
	_ = v.RegisterValidation("passwd", func(fl validator.FieldLevel) bool {
		return len(fl.Field().String()) > 6
	})

	err := v.Struct(user)

	if err != nil {
		return errors.NewBadRequestError(err.Error())
	}
	return nil
}
