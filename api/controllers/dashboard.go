package controllers

import (
	"server/api/middlewares"
	"server/api/services"
	"server/api/utils"

	"github.com/gin-gonic/gin"
)

func RegisterDashboardRoutes(app *gin.RouterGroup) {
	router := app.Group("dashboard", middlewares.Auth())

	router.GET("", func(c *gin.Context) {
		keyCnt, lockCnt, personelCnt := services.DashboardData()
		utils.MakeResponseSuccess(c, gin.H{
			"key_cnt":      keyCnt,
			"lock_cnt":     lockCnt,
			"personel_cnt": personelCnt,
		}, nil)
	})
}
