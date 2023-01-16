package controllers

import (
	"server/api/utils"

	"github.com/gin-gonic/gin"
)

func RegisterLog(app *gin.Engine) {
	router := app.Group("/log")

	router.GET("/access", func(c *gin.Context) {
		utils.ResponseUnimplemented(c)
	})

	router.GET("/healthcheck", func(c *gin.Context) {
		utils.ResponseUnimplemented(c)
	})

	router.GET("/rssi", func(c *gin.Context) {
		utils.ResponseUnimplemented(c)
	})
}
