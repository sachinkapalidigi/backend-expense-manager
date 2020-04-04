package services

import (
	"github.com/sachinkapalidigi/backend-expense-manager/domain/users"
	"github.com/sachinkapalidigi/backend-expense-manager/utils/accesstokenutil"
	"github.com/sachinkapalidigi/backend-expense-manager/utils/cryptoutil"
	"github.com/sachinkapalidigi/backend-expense-manager/utils/dateutil"
	"github.com/sachinkapalidigi/backend-expense-manager/utils/errors"
)

var (
	UsersServices usersServicesInterface = &usersServices{}
)

type usersServices struct{}

type usersServicesInterface interface {
	Create(users.User) (*users.User, *errors.RestErr)
	Authenticate(*users.User) (*accesstokenutil.AccessToken, *errors.RestErr)
}

func (s *usersServices) Create(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	user.CreatedAt = dateutil.GetNowDBFormat()
	user.Password = cryptoutil.GetMd5(user.Password)

	if err := user.CreateUser(); err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *usersServices) Authenticate(currentUser *users.User) (*accesstokenutil.AccessToken, *errors.RestErr) {
	if err := currentUser.ValidateLogin(); err != nil {
		return nil, err
	}
	currentUser.Password = cryptoutil.GetMd5(currentUser.Password)
	var user = users.User{
		Email: currentUser.Email,
	}
	if err := user.GetPasswordByEmail(); err != nil {
		return nil, err
	}

	if currentUser.Password != user.Password {
		return nil, errors.NewBadRequestError("Invalid Email Id or Password")
	}
	// generate token
	accessToken, err := accesstokenutil.GenerateToken(user.ID)
	if err != nil {
		return nil, err
	}

	return accessToken, nil

}
