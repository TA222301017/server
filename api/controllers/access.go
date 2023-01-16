package controllers

import (
	"server/api/utils"

	"github.com/gin-gonic/gin"
)

func RegisterAccess(app *gin.Engine) {
	router := app.Group("/access")

	router.GET("/:personel_id", func(c *gin.Context) {
		utils.ResponseUnimplemented(c)
	})

	router.POST("/", func(c *gin.Context) {
		utils.ResponseUnimplemented(c)
	})

	router.PATCH("/:access_rule_id", func(c *gin.Context) {
		utils.ResponseUnimplemented(c)
	})

	router.DELETE("/:access_rule_id", func(c *gin.Context) {
		utils.ResponseUnimplemented(c)
	})
}
