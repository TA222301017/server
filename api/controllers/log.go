package controllers

import (
	"errors"
	"fmt"
	"server/api/services"
	"server/api/utils"
	"strconv"

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
		temp := c.Query("personel_id")
		personelID, err := strconv.ParseUint(temp, 10, 64)
		fmt.Println(temp, personelID)
		if err != nil {
			utils.ResponseBadRequest(c, errors.New("invalid personel_id"))
			return
		}

		params := utils.ParseSearchParameter(c)

		data, pagination, err := services.GetRSSILog(params, personelID)
		if err != nil {
			utils.ResponseServerError(c, err)
			return
		}

		utils.MakeResponseSuccess(c, data, pagination)
	})
}
