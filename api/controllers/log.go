package controllers

import (
	"server/api/services"
	"server/api/utils"

	"github.com/gin-gonic/gin"
)

func RegisterLog(app *gin.Engine) {
	router := app.Group("/log")

	router.GET("/access", func(c *gin.Context) {
		params := utils.ParseSearchParameter(c)
		location := c.Query("location")
		personel := c.Query("personel")

		data, pagination, err := services.GetAccessLog(*params, location, personel)
		if err != nil {
			utils.ResponseServerError(c, err)
			return
		}

		utils.MakeResponseSuccess(c, data, pagination)
	})

	router.GET("/healthcheck", func(c *gin.Context) {
		params := utils.ParseSearchParameter(c)

		data, pagination, err := services.GetHealthcheckLog(params)
		if err != nil {
			utils.ResponseServerError(c, err)
			return
		}

		utils.MakeResponseSuccess(c, data, pagination)
	})

	router.GET("/rssi", func(c *gin.Context) {
		keyword := c.Query("keyword")
		params := utils.ParseSearchParameter(c)

		data, pagination, err := services.GetRSSILog(params, keyword)
		if err != nil {
			utils.ResponseServerError(c, err)
			return
		}

		utils.MakeResponseSuccess(c, data, pagination)
	})
}
