package controllers

import (
	"server/api/utils"

	"github.com/gin-gonic/gin"
)

func RegisterHello(app *gin.Engine) {
	router := app.Group("/hello")

	router.GET("/", func(c *gin.Context) {
		utils.MakeResponseSuccess(c, "halo dunia", nil, "")
	})
}
