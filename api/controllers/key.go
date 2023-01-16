package controllers

import (
	"server/api/utils"

	"github.com/gin-gonic/gin"
)

func RegisterKey(app *gin.Engine) {
	router := app.Group("/device/key")

	router.GET("/", func(c *gin.Context) {
		utils.ResponseUnimplemented(c)
	})

	router.GET("/:key_id", func(c *gin.Context) {
		utils.ResponseUnimplemented(c)
	})

	router.POST("/", func(c *gin.Context) {
		utils.ResponseUnimplemented(c)
	})

	router.PATCH("/", func(c *gin.Context) {
		utils.ResponseUnimplemented(c)
	})
}
