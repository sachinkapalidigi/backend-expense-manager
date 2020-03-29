package app

import (
	"github.com/gin-gonic/gin"
	"github.com/sachinkapalidigi/backend-expense-manager/logger"
)

var (
	router = gin.Default()
)

// StartApplication : starts go lang application
func StartApplication() {
	mapUrls()

	// start logger
	logger.Info("Starting Server")
	router.Run(":8080")
}
