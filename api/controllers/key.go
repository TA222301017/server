package controllers

import (
	"server/api/services"
	"server/api/template"
	"server/api/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RegisterKey(app *gin.Engine) {
	router := app.Group("/device/key")

	router.GET("", func(c *gin.Context) {
		params := utils.ParseSearchParameter(c)
		keyword := c.Query("keyword")
		status := c.Query("status")

		data, pagination, err := services.GetKeys(*params, keyword, status)
		if err != nil {
			utils.ResponseServerError(c, err)
			return
		}

		utils.MakeResponseSuccess(c, data, pagination)
	})

	router.GET("/:key_id", func(c *gin.Context) {
		temp := c.Param("key_id")
		keyID, err := strconv.ParseUint(temp, 10, 64)
		if err != nil {
			utils.ResponseBadRequest(c, err)
			return
		}

		data, err := services.GetKey(keyID)
		if err != nil {
			utils.ResponseServerError(c, err)
			return
		}

		utils.MakeResponseSuccess(c, data, nil)
	})

	router.POST("", func(c *gin.Context) {
		var body template.AddKeyRequest
		if err := c.Bind(&body); err != nil {
			utils.ResponseServerError(c, err)
			return
		}

		if err := body.Validate(); err != nil {
			utils.ResponseBadRequest(c, err)
			return
		}

		data, err := services.AddKey(body)
		if err != nil {
			utils.ResponseServerError(c, err)
			return
		}

		utils.MakeResponseCreated(c, data)
	})

	router.PATCH("/:key_id", func(c *gin.Context) {
		temp := c.Param("key_id")
		keyID, err := strconv.ParseUint(temp, 10, 64)
		if err != nil {
			utils.ResponseBadRequest(c, err)
			return
		}

		var body template.EditKeyRequest
		if err := c.Bind(&body); err != nil {
			utils.ResponseServerError(c, err)
			return
		}

		if err := body.Validate(); err != nil {
			utils.ResponseBadRequest(c, err)
			return
		}

		data, err := services.EditKey(body, keyID)
		if err != nil {
			utils.ResponseServerError(c, err)
			return
		}

		utils.MakeResponseSuccess(c, data, nil)
	})
}
