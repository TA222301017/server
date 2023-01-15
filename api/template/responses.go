package template

type Pagination struct {
	Page  int `json:"page,omitempty"`
	Limit int `json:"limit,omitempty"`
	Last  int `json:"last,omitempty"`
	Total int `json:"total,omitempty"`
}

type BaseResponse struct {
	Data       interface{} `json:"data,omitempty"`
	Msg        string      `json:"msg,omitempty"`
	Error      string      `json:"error,omitempty"`
	Token      string      `json:"token,omitempty"`
	Pagination *Pagination `json:"pagination,omitempty"`
}
