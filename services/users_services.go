package services

import (
	"github.com/sachinkapalidigi/backend-expense-manager/domain/users"
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
