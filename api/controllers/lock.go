package controllers

import (
	"server/api/middlewares"
	"server/api/services"
	"server/api/template"
	"server/api/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RegisterLockRoutes(app *gin.RouterGroup) {
	router := app.Group("/device/lock", middlewares.Auth())

	router.GET("", func(c *gin.Context) {
		params := utils.ParseSearchParameter(c)
		keyword := c.Query("keyword")
		status := c.Query("status")

		data, pagination, err := services.GetLocks(*params, keyword, status)
		if err != nil {
			utils.ResponseServerError(c, err)
			return
		}

		utils.MakeResponseSuccess(c, data, pagination)
	})

	router.GET("/:lock_id", func(c *gin.Context) {
		temp := c.Param("lock_id")
		lockID, err := strconv.ParseUint(temp, 10, 64)
		if err != nil {
			utils.ResponseBadRequest(c, err)
			return
		}

		data, err := services.GetLock(lockID)
		if err != nil {
			utils.ResponseServerError(c, err)
			return
		}

		utils.MakeResponseSuccess(c, data, nil)
	})

	router.PATCH("/:lock_id", func(c *gin.Context) {
		temp := c.Param("lock_id")
		lockID, err := strconv.ParseUint(temp, 10, 64)
		if err != nil {
			utils.ResponseBadRequest(c, err)
			return
		}

		var body template.EditLockRequest
		if err := c.Bind(&body); err != nil {
			utils.ResponseServerError(c, err)
			return
		}

		data, err := services.EditLock(body, lockID)
		if err != nil {
			utils.ResponseServerError(c, err)
			return
		}

		utils.MakeResponseSuccess(c, data, nil)
	})

	router.GET("/check", func(c *gin.Context) {
		data, err := services.CheckLocks()
		if err != nil {
			utils.ResponseServerError(c, err)
			return
		}

		utils.MakeResponseSuccess(c, data, nil)
	})

	router.GET("/check/:lock_id", func(c *gin.Context) {
		temp := c.Param("lock_id")
		lockID, err := strconv.ParseUint(temp, 10, 64)
		if err != nil {
			utils.ResponseBadRequest(c, err)
			return
		}

		data, err := services.CheckLock(lockID)
		if err != nil {
			utils.ResponseServerError(c, err)
			return
		}

		utils.MakeResponseSuccess(c, data, nil)
	})
}
