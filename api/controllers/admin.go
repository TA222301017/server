package controllers

import (
	"server/api/services"
	"server/api/template"
	"server/api/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RegisterAdminRoutes(app *gin.RouterGroup) {
	r := app.Group("/admin")

	r.POST("", func(c *gin.Context) {
		var body template.CreateAdminRequest
		if err := c.BindJSON(&body); err != nil {
			utils.ResponseServerError(c, err)
			return
		}

		if err := body.Validate(); err != nil {
			utils.ResponseBadRequest(c, err)
			return
		}

		admin, err := services.CreateAdmin(body)
		if err != nil {
			utils.ResponseServerError(c, err)
			return
		}

		utils.MakeResponseCreated(c, admin)
	})

	r.GET("", func(c *gin.Context) {
		p := utils.ParseSearchParameter(c)
		keyword := c.Query("keyword")

		data, pagination, err := services.GetAdmins(*p, keyword)
		if err != nil {
			utils.ResponseServerError(c, err)
			return
		}

		utils.MakeResponseSuccess(c, data, pagination)
	})

	r.GET("/:admin_id", func(c *gin.Context) {
		temp := c.Param("admin_id")
		adminID, err := strconv.ParseUint(temp, 10, 64)
		if err != nil {
			utils.ResponseServerError(c, err)
			return
		}

		data, err := services.GetAdmin(adminID)
		if err != nil {
			utils.ResponseServerError(c, err)
			return
		}

		utils.MakeResponseSuccess(c, data, nil)
	})

	r.PATCH("/:admin_id", func(c *gin.Context) {
		temp := c.Param("admin_id")
		adminID, err := strconv.ParseUint(temp, 10, 64)
		if err != nil {
			utils.ResponseServerError(c, err)
			return
		}

		var body template.EditAdminRequest
		if err := c.BindJSON(&body); err != nil {
			utils.ResponseServerError(c, err)
			return
		}

		admin, err := services.EditAdmin(adminID, body)
		if err != nil {
			utils.ResponseServerError(c, err)
			return
		}

		utils.MakeResponseSuccess(c, admin, nil)
	})

	r.DELETE("/:admin_id", func(c *gin.Context) {
		temp := c.Param("admin_id")
		adminID, err := strconv.ParseUint(temp, 10, 64)
		if err != nil {
			utils.ResponseServerError(c, err)
			return
		}

		if err := services.DeleteAdmin(adminID); err != nil {
			utils.ResponseServerError(c, err)
			return
		}

		utils.MakeResponseSuccess(c, nil, nil)
	})
}
