package api

import (
	"server/api/controllers"
	"server/api/middlewares"
	"server/api/setup"

	"github.com/gin-gonic/gin"
)

func RegisterAPIRoutes(app *gin.Engine) *gin.RouterGroup {
	setup.Mode()

	api := app.Group("/api")
	// setup.CustomErrorHandler(api)
	setup.CustomLogger(api)

	controllers.RegisterAuthRoutes(api)

	api.Use(middlewares.Auth())

	controllers.RegisterDashboardRoutes(api)
	controllers.RegisterAccessRoutes(api)
	controllers.ResgisterPersonelRoutes(api)
	controllers.RegisterLogRoutes(api)
	controllers.RegisterKeyRoutes(api)
	controllers.RegisterLockRoutes(api)
	controllers.RegisterAdminRoutes(api)
	controllers.ResgiterPlanRoutes(api)

	return api
}
