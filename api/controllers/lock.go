package controllers

import (
	"server/api/utils"

	"github.com/gin-gonic/gin"
)

func RegisterLock(app *gin.Engine) {
	router := app.Group("/device/lock")

	router.GET("/", func(c *gin.Context) {
		utils.ResponseUnimplemented(c)
	})

	router.GET("/:lock_id", func(c *gin.Context) {
		utils.ResponseUnimplemented(c)
	})

	router.PATCH("/:lock_id", func(c *gin.Context) {
		utils.ResponseUnimplemented(c)
	})

	router.GET("/check", func(c *gin.Context) {
		utils.ResponseUnimplemented(c)
	})

	router.GET("/check/:lock_id", func(c *gin.Context) {
		utils.ResponseUnimplemented(c)
	})
}
