package template

type Pagination struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
	Last  int `json:"last"`
	Total int `json:"total"`
}

type BaseResponse struct {
	Data       interface{} `json:"data,omitempty"`
	Msg        string      `json:"msg,omitempty"`
	Error      string      `json:"error,omitempty"`
	Token      string      `json:"token,omitempty"`
	Pagination *Pagination `json:"pagination,omitempty"`
}
