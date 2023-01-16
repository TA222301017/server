package models

type Key struct {
	BaseModel
	KeyID       uint64 `json:"key_id" gorm:"unique"`
	Label       string `json:"label"`
	Status      bool   `json:"status"`
	Description string `json:"description"`
}
