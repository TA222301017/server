package web

import (
	"server/api"
	"server/web/setup"

	"github.com/gin-gonic/gin"
)

func Run() {
	app := gin.New()
	address := setup.GetAddress()

	setup.Cors(app)
	setup.Static(app)

	api.RegisterAPIRoutes(app)

	app.NoRoute(func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	app.Run(address)
}
