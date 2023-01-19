package utils

import (
	"server/api/template"
	"time"

	"github.com/gin-gonic/gin"
)

func ParseSearchParameter(c *gin.Context) *template.SearchParameter {
	params := template.SearchParameter{}
	if err := c.BindQuery(&params); err != nil {
		return nil
	}

	if params.Page == 0 {
		params.Page = 1
	}

	if params.Limit == 0 {
		params.Limit = 20
	}

	if params.StartDate.IsZero() {
		params.StartDate = time.Now()
	}

	if params.EndDate.IsZero() {
		params.EndDate = time.Now()
	}

	return &params
}
