package controllers

import (
	"server/api/services"
	"server/api/utils"

	"github.com/gin-gonic/gin"
)

func RegisterDashboard(app *gin.Engine) {
	router := app.Group("dashboard")

	router.GET("/", func(c *gin.Context) {
		keyCnt, lockCnt, personelCnt := services.DashboardData()
		utils.MakeResponseSuccess(c, gin.H{
			"key_cnt":      keyCnt,
			"lock_cnt":     lockCnt,
			"personel_cnt": personelCnt,
		}, nil)
	})
}
