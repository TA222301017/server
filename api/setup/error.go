package setup

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CustomErrorHandler(app *gin.Engine) {
	app.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{
			"msg": "page not found",
		})
	})

	app.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, gin.H{
			"msg": "method not allowed",
		})
	})
}
