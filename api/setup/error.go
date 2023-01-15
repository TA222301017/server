package setup

import (
	"net/http"
	"strings"

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

	app.Use(func(c *gin.Context) {
		defer func(c *gin.Context) {
			c.Next()

			c.JSON(http.StatusInternalServerError, gin.H{
				"msg": strings.Join(c.Errors.Errors(), " ; "),
			})
		}(c)
	})
}
