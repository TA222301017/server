package controllers

import (
	"errors"
	"server/api/services"
	"server/api/template"
	"server/api/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RegisterAccess(app *gin.Engine) {
	router := app.Group("/access")

	router.GET("/:personel_id", func(c *gin.Context) {
		temp := c.Param("personel_id")
		personelID, err := strconv.ParseUint(temp, 10, 64)
		if err != nil {
			utils.ResponseBadRequest(c, errors.New("invalid personel_id"))
			return
		}

		data := services.GetPersonelAccessRules(personelID)
		utils.MakeResponseSuccess(c, data, nil)
	})

	router.POST("/", func(c *gin.Context) {
		var body template.AddAccessRule
		if err := c.Bind(&body); err != nil {
			utils.ResponseServerError(c, err)
			return
		}

		claims, err := utils.GetClaimsFromContext(c)
		if err != nil {
			utils.ResponseUnauthorized(c, err)
			return
		}

		data, err := services.AddAccessRule(body, claims.ID)
		if err != nil {
			utils.ResponseServerError(c, err)
			return
		}

		utils.MakeResponseCreated(c, data)
	})

	router.PATCH("/:access_rule_id", func(c *gin.Context) {
		var body template.EditAccessRule
		if err := c.Bind(&body); err != nil {
			utils.ResponseServerError(c, err)
			return
		}

		temp := c.Param("access_rule_id")
		accessRuleID, err := strconv.ParseUint(temp, 10, 64)
		if err != nil {
			utils.ResponseBadRequest(c, errors.New("invalid access_rule_id"))
			return
		}

		claims, err := utils.GetClaimsFromContext(c)
		if err != nil {
			utils.ResponseUnauthorized(c, err)
			return
		}

		data, err := services.EditAccessRule(body, claims.ID, accessRuleID)
		if err != nil {
			utils.ResponseServerError(c, err)
			return
		}

		utils.MakeResponseSuccess(c, data, nil)
	})

	router.DELETE("/:access_rule_id", func(c *gin.Context) {
		temp := c.Param("access_rule_id")
		accessRuleID, err := strconv.ParseUint(temp, 10, 64)
		if err != nil {
			utils.ResponseBadRequest(c, errors.New("invalid access_rule_id"))
			return
		}

		if err := services.DeleteAccessRule(accessRuleID); err != nil {
			utils.ResponseServerError(c, err)
			return
		}

		utils.MakeResponseSuccess(c, nil, nil)
	})
}
