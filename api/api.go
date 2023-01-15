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

	setup.CustomErrorHandler(api)
	setup.CustomLogger(api)
	setup.Cors(api)
	setup.Mode()

	controllers.RegisterHello(api)

	api.Use(middlewares.Auth())

	api.Run(address)
}
