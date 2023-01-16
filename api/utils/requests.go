package utils

import (
	"server/api/template"

	"github.com/gin-gonic/gin"
)

func ParseSearchParameter(c *gin.Context) *template.SearchParameter {
	params := template.SearchParameter{}
	if err := c.Bind(&params); err == nil {
		return &params
	}
	return nil
}
