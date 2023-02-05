package api

import (
	"server/api/controllers"
	"server/api/middlewares"
	"server/api/setup"

	"github.com/gin-gonic/gin"
)

func Run() {
	api := gin.New()
	address := setup.GetAddress()

	// setup.CustomErrorHandler(api)
	setup.CustomLogger(api)
	setup.Cors(api)
	setup.Mode()

	controllers.RegisterAuthRoutes(api)

	api.Use(middlewares.Auth())

	controllers.RegisterDashboardRoutes(api)
	controllers.RegisterAccessRoutes(api)
	controllers.ResgisterPersonelRoutes(api)
	controllers.RegisterLogRoutes(api)
	controllers.RegisterKeyRoutes(api)
	controllers.RegisterLockRoutes(api)

	api.Run(address)
}
