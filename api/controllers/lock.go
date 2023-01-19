package controllers

import (
	"server/api/services"
	"server/api/template"
	"server/api/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RegisterLock(app *gin.Engine) {
	router := app.Group("/device/lock")

	router.GET("/", func(c *gin.Context) {
		params := utils.ParseSearchParameter(c)
		keyword := c.Query("keyword")
		status := c.GetBool("status")

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
		utils.ResponseUnimplemented(c)
	})

	router.GET("/check/:lock_id", func(c *gin.Context) {
		utils.ResponseUnimplemented(c)
	})
}
