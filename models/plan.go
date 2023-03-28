package models

type Plan struct {
	BaseModel
	Name        string  `json:"name" gorm:"unique"`
	Width       float32 `json:"width"`
	Height      float32 `json:"height"`
	ImageWidth  int     `json:"image_width"`
	ImageHeight int     `json:"image_height"`
	ImageURL    string  `json:"image_url"`
}
