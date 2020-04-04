package app

import (
	"github.com/gin-gonic/gin"
	"github.com/sachinkapalidigi/backend-expense-manager/logger"
	"github.com/sachinkapalidigi/backend-expense-manager/middlewares"
)

var (
	router = gin.Default()
)

// StartApplication : starts go lang application
func StartApplication() {

	router.Use(middlewares.AuthMiddleware.UserLoader())
	mapUrls()

	// start logger
	logger.Info("Starting Server")
	router.Run(":8080")
}
