package controllers

import (
	"errors"
	"server/api/services"
	"server/api/template"
	"server/api/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(app *gin.Engine) {
	app.POST("/login", func(c *gin.Context) {
		var body template.LoginRequest
		if err := c.Bind(&body); err != nil {
			utils.ResponseServerError(c, err)
			return
		}

		if err := body.Validate(); err != nil {
			utils.ResponseBadRequest(c, err)
			return
		}

		user, token, err := services.Login(body)
		if err != nil {
			utils.ResponseServerError(c, err)
			return
		}

		c.Set("token", token)
		utils.MakeResponseSuccess(c, user, nil, "login successful")
	})

	app.POST("/register", func(c *gin.Context) {
		var body template.RegisterRequest
		if err := c.Bind(&body); err != nil {
			utils.ResponseServerError(c, err)
			return
		}

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			utils.ResponseUnauthorized(c, errors.New("authorization header is not found"))
			return
		}

		temp := strings.Split(authHeader, " ")
		if len(temp) != 2 {
			utils.ResponseUnauthorized(c, errors.New("authorization header must use the format Bearer TOKEN"))
			return
		}

		if _, err := utils.VerifyJWT(temp[1]); err != nil {
			utils.ResponseUnauthorized(c, err)
			return
		}

		if err := body.Validate(); err != nil {
			utils.ResponseBadRequest(c, err)
			return
		}

		user, err := services.RegisterNewUser(body)
		if err != nil {
			utils.ResponseServerError(c, err)
			return
		}

		utils.MakeResponseCreated(c, user, "new user created successfully")
	})
}
