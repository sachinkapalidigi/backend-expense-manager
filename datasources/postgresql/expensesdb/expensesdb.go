package expensesdb

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/sachinkapalidigi/backend-expense-manager/logger"

	_ "github.com/lib/pq"
)

const (
	postgressql_username    = "postgressql_username"
	postgressql_password    = "postgressql_password"
	postgressql_host        = "postgressql_host"
	postgresql_port         = "postgresql_port"
	postgressql_expenses_db = "postgressql_expenses_db"
)

var (
	Client *sql.DB

	username = os.Getenv(postgressql_username)
	password = os.Getenv(postgressql_password)
	host     = os.Getenv(postgressql_host)
	schema   = os.Getenv(postgressql_expenses_db)
	port     = os.Getenv(postgresql_port)
)

func init() {
	dbSourceName := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable", host, port, username, password, schema)

	var err error

	Client, err = sql.Open("postgres", dbSourceName)
	if err != nil {
		logger.Error("Error while connecting to db", err)
		panic(err)
	}

	if err = Client.Ping(); err != nil {
		panic(err)
	}
	logger.Info("Database successfully configured")
}
