package template

import "time"

type SearchParameter struct {
	Page      int       `json:"page" form:"page"`
	Limit     int       `json:"limit" form:"limit"`
	StartDate time.Time `json:"startdate" form:"startdate"`
	EndDate   time.Time `json:"enddate" form:"enddate"`
}
