package controllers

import (
	"server/api/middlewares"
	"server/api/services"
	"server/api/template"
	"server/api/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ResgisterPersonelRoutes(app *gin.RouterGroup) {
	router := app.Group("/personel", middlewares.Auth())

	router.GET("/role", func(c *gin.Context) {
		data := services.GetRoles()

		utils.MakeResponseSuccess(c, data, nil)
	})

	router.GET("", func(c *gin.Context) {
		params := utils.ParseSearchParameter(c)
		keyword := c.Query("keyword")
		status := c.Query("status")

		data, pagination, err := services.GetPersonels(*params, keyword, status)
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

	router.POST("", func(c *gin.Context) {
		var body template.AddPersonelRequest
		if err := c.Bind(&body); err != nil {
			utils.ResponseServerError(c, err)
			return
		}

		if err := body.Validate(); err != nil {
			utils.ResponseBadRequest(c, err)
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
		temp := c.Param("personel_id")
		personelID, err := strconv.ParseUint(temp, 10, 64)
		if err != nil {
			utils.ResponseBadRequest(c, err)
			return
		}

		var body template.EditPersonelRequest
		if err := c.Bind(&body); err != nil {
			utils.ResponseServerError(c, err)
			return
		}

		data, err := services.EditPersonel(body, personelID)
		if err != nil {
			utils.ResponseServerError(c, err)
			return
		}

		utils.MakeResponseSuccess(c, data, nil)
	})
}
