package models

type Key struct {
	BaseModel
	KeyID       string `json:"key_id" gorm:"unique"`
	Label       string `json:"name"`
	Status      bool   `json:"status"`
	AESKey      string `json:"aes_key,omitempty"`
	Description string `json:"description"`
}
