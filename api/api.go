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

	controllers.RegisterAuth(api)

	api.Use(middlewares.Auth())

	controllers.RegisterDashboard(api)
	controllers.RegisterAccess(api)
	controllers.ResgisterPersonel(api)
	controllers.RegisterLog(api)
	controllers.RegisterKey(api)
	controllers.RegisterLock(api)

	api.Run(address)
}
