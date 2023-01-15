package utils

import (
	"net/http"
	"server/api/template"

	"github.com/gin-gonic/gin"
)

func MakeResponse(c *gin.Context, status int, data interface{}, pagination *template.Pagination, msg string, err string, token string) {
	c.JSON(status, template.BaseResponse{
		Data:       data,
		Pagination: pagination,
		Msg:        msg,
		Error:      err,
		Token:      token,
	})
}

func MakeResponseSuccess(c *gin.Context, data interface{}, pagination *template.Pagination, msg ...string) {
	message := "ok"
	if len(msg) != 0 {
		message = msg[0]
	}

	v, exists := c.Get("token")
	token, ok := v.(string)
	if !exists || !ok {
		token = ""
	}

	c.JSON(http.StatusOK, template.BaseResponse{
		Data:       data,
		Pagination: pagination,
		Msg:        message,
		Error:      "",
		Token:      token,
	})
}

func MakeResponseCreated(c *gin.Context, data interface{}, msg ...string) {
	message := "ok"
	if len(msg) != 0 {
		message = msg[0]
	}

	v, exists := c.Get("token")
	token, ok := v.(string)
	if !exists || !ok {
		token = ""
	}

	c.JSON(http.StatusOK, template.BaseResponse{
		Data:       data,
		Pagination: nil,
		Msg:        message,
		Error:      "",
		Token:      token,
	})
}

func MakeResponseError(c *gin.Context, status int, msg string, err string) {

	v, exists := c.Get("token")
	token, ok := v.(string)
	if !exists || !ok {
		token = ""
	}

	c.JSON(http.StatusOK, template.BaseResponse{
		Data:       nil,
		Pagination: nil,
		Msg:        msg,
		Error:      err,
		Token:      token,
	})
}
