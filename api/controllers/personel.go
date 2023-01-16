package controllers

import (
	"server/api/utils"

	"github.com/gin-gonic/gin"
)

func ResgisterPersonel(app *gin.Engine) {
	router := app.Group("/personel")

	router.GET("/", func(c *gin.Context) {
		utils.ResponseUnimplemented(c)
	})

	router.GET("/:personel_id", func(c *gin.Context) {
		utils.ResponseUnimplemented(c)
	})

	router.POST("/", func(c *gin.Context) {
		utils.ResponseUnimplemented(c)
	})

	router.PATCH("/:personel_id", func(c *gin.Context) {
		utils.ResponseUnimplemented(c)
	})
}
