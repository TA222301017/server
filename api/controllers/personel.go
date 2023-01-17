package controllers

import (
	"server/api/services"
	"server/api/template"
	"server/api/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ResgisterPersonel(app *gin.Engine) {
	router := app.Group("/personel")

	router.GET("/", func(c *gin.Context) {
		params := utils.ParseSearchParameter(c)
		keyword := c.GetString("keyword")
		status := c.GetBool("status")

		data, pagination, err := services.GetPersonels(*params, status, keyword)
		if err != nil {
			utils.ResponseServerError(c, err)
			return
		}

		utils.MakeResponseSuccess(c, data, pagination)
	})

	router.GET("/:personel_id", func(c *gin.Context) {
		temp := c.Param("personel_id")
		personelID, err := strconv.ParseUint(temp, 10, 64)
		if err != nil {
			utils.ResponseBadRequest(c, err)
			return
		}

		data, err := services.GetPersonel(personelID)
		if err != nil {
			utils.ResponseServerError(c, err)
			return
		}

		utils.MakeResponseSuccess(c, data, nil)
	})

	router.POST("/", func(c *gin.Context) {
		var body template.AddPersonelRequest
		if err := c.Bind(&body); err != nil {
			utils.ResponseServerError(c, err)
			return
		}

		data, err := services.RegisterPersonel(body)
		if err != nil {
			utils.ResponseServerError(c, err)
			return
		}

		utils.MakeResponseCreated(c, data)
	})

	router.PATCH("/:personel_id", func(c *gin.Context) {
		utils.ResponseUnimplemented(c)
	})
}
