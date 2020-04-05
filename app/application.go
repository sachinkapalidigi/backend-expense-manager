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

	// config := cors.DefaultConfig()
	// config.AllowHeaders = []string{"Origin", "Authorization"}
	// config.AllowAllOrigins = true
	// router.Use(cors.Default())
	router.Use(middlewares.CorsMiddleware.CORSMiddleware())
	router.Use(middlewares.AuthMiddleware.UserLoader())
	mapUrls()

	// start logger
	logger.Info("Starting Server")
	router.Run(":8080")
}
