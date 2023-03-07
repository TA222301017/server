package models

type Plan struct {
	BaseModel
	Name     string  `json:"name"`
	Width    float32 `json:"width"`
	Height   float32 `json:"height"`
	ImageURL string  `json:"image_url"`
}
