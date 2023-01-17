package models

type Key struct {
	BaseModel
	KeyID       string `json:"key_id" gorm:"unique"`
	Label       string `json:"name"`
	Status      bool   `json:"status"`
	Description string `json:"description"`
}
