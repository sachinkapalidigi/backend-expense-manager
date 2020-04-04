package users

import (
	"github.com/sachinkapalidigi/backend-expense-manager/datasources/postgresql/expensesdb"
	"github.com/sachinkapalidigi/backend-expense-manager/logger"
	"github.com/sachinkapalidigi/backend-expense-manager/utils/errors"
)

const (
	queryInsertUser        = "INSERT INTO users (name, email, password, created_at) VALUES ($1, $2, $3, $4) RETURNING id;"
	querySelectUserByEmail = "SELECT id, name, email, password, created_at FROM users WHERE email=$1;"
)

// CreateUser : add new user to database
func (user *User) CreateUser() *errors.RestErr {
	stmt, err := expensesdb.Client.Prepare(queryInsertUser)
	if err != nil {
		logger.Error("Prepare statement error", err)
		return errors.NewInternalServerError("Couldnot create User")
	}
	defer stmt.Close()

	if err := stmt.QueryRow(&user.Name, &user.Email, &user.Password, &user.CreatedAt).Scan(&user.ID); err != nil {
		logger.Error("Errors in saving / retrieving user id", err)
		return errors.NewInternalServerError("Couldnot Save User")
	}

	return nil
}

// GetPasswordByEmail : Get the password/user details from Email address
func (user *User) GetPasswordByEmail() *errors.RestErr {
	stmt, err := expensesdb.Client.Prepare(querySelectUserByEmail)
	if err != nil {
		logger.Error("Error in preparing statement", err)
		return errors.NewInternalServerError("Couldnot Get User details")
	}

	defer stmt.Close()

	row := stmt.QueryRow(user.Email)
	if err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt); err != nil {

		return errors.NewBadRequestError("Invalid Email Address")
	}

	return nil
}
