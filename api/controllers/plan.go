package controllers

import (
	"server/api/middlewares"
	"server/api/services"
	"server/api/template"
	"server/api/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ResgiterPlanRoutes(app *gin.RouterGroup) {
	r := app.Group("/plan", middlewares.Auth())

	r.GET("", func(c *gin.Context) {
		keyword := c.Query("keyword")
		p := utils.ParseSearchParameter(c)

		data, pagination, err := services.GetPlans(*p, keyword)
		if err != nil {
			utils.ResponseServerError(c, err)
			return
		}

		utils.MakeResponseSuccess(c, data, pagination)
	})

	r.GET("/:plan_id", func(c *gin.Context) {
		temp := c.Param("plan_id")
		planID, err := strconv.ParseUint(temp, 10, 64)
		if err != nil {
			utils.ResponseBadRequest(c, err)
			return
		}

		data, err := services.GetPlan(planID)
		if err != nil {
			utils.ResponseServerError(c, err)
			return
		}

		utils.MakeResponseSuccess(c, data, nil)
	})

	r.POST("", func(c *gin.Context) {
		var body template.CreatePlanRequest
		if err := c.Bind(&body); err != nil {
			utils.ResponseServerError(c, err)
			return
		}

		if err := body.Validate(); err != nil {
			utils.ResponseBadRequest(c, err)
			return
		}

		data, err := services.CreatePlan(body)
		if err != nil {
			utils.ResponseServerError(c, err)
			return
		}

		utils.MakeResponseCreated(c, data)
	})

	r.PATCH("/:plan_id", func(c *gin.Context) {
		temp := c.Param("plan_id")
		planID, err := strconv.ParseUint(temp, 10, 64)
		if err != nil {
			utils.ResponseBadRequest(c, err)
			return
		}

		var body template.CreatePlanRequest
		if err := c.Bind(&body); err != nil {
			utils.ResponseServerError(c, err)
			return
		}

		data, err := services.EditPlan(body, planID)
		if err != nil {
			utils.ResponseServerError(c, err)
			return
		}

		utils.MakeResponseSuccess(c, data, nil)
	})

	r.DELETE("/:plan_id", func(c *gin.Context) {
		temp := c.Param("plan_id")
		planID, err := strconv.ParseUint(temp, 10, 64)
		if err != nil {
			utils.ResponseBadRequest(c, err)
			return
		}

		if err := services.DeletePlan(planID); err != nil {
			utils.ResponseServerError(c, err)
			return
		}

		utils.MakeResponseSuccess(c, "ok", nil)
	})

	r.POST("/:plan_id/lock", func(c *gin.Context) {
		temp := c.Param("plan_id")
		planID, err := strconv.ParseUint(temp, 10, 64)
		if err != nil {
			utils.ResponseBadRequest(c, err)
			return
		}

		var body template.AddLockToPlanRequest
		if err := c.Bind(&body); err != nil {
			utils.ResponseServerError(c, err)
			return
		}

		if err := body.Validate(); err != nil {
			utils.ResponseBadRequest(c, err)
			return
		}

		if err := services.AddLockToPlan(body, planID); err != nil {
			utils.ResponseServerError(c, err)
			return
		}

		data, err := services.GetPlan(planID)
		if err != nil {
			utils.ResponseServerError(c, err)
			return
		}

		utils.MakeResponseSuccess(c, data, nil)
	})

	r.PATCH("/:plan_id/lock/:lock_id", func(c *gin.Context) {
		temp := c.Param("plan_id")
		planID, err := strconv.ParseUint(temp, 10, 64)
		if err != nil {
			utils.ResponseBadRequest(c, err)
			return
		}

		temp = c.Param("lock_id")
		lockID, err := strconv.ParseUint(temp, 10, 64)
		if err != nil {
			utils.ResponseBadRequest(c, err)
			return
		}

		var body template.AddLockToPlanRequest
		if err := c.Bind(&body); err != nil {
			utils.ResponseServerError(c, err)
			return
		}

		if err := services.EditLockInPlan(body, planID, lockID); err != nil {
			utils.ResponseServerError(c, err)
			return
		}

		data, err := services.GetPlan(planID)
		if err != nil {
			utils.ResponseServerError(c, err)
			return
		}

		utils.MakeResponseSuccess(c, data, nil)
	})

	r.DELETE("/:plan_id/lock/:lock_id", func(c *gin.Context) {
		temp := c.Param("plan_id")
		planID, err := strconv.ParseUint(temp, 10, 64)
		if err != nil {
			utils.ResponseBadRequest(c, err)
			return
		}

		temp = c.Param("lock_id")
		lockID, err := strconv.ParseUint(temp, 10, 64)
		if err != nil {
			utils.ResponseBadRequest(c, err)
			return
		}

		if err := services.RemoveLockFromPlan(planID, lockID); err != nil {
			utils.ResponseServerError(c, err)
			return
		}

		utils.MakeResponseSuccess(c, "ok", nil)
	})
}
