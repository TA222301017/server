package models

type Lock struct {
	BaseModel
	LockID      uint64 `json:"lock_id" gorm:"unique"`
	IpAddress   string `json:"ip_address"`
	Label       string `json:"label"`
	Description string `json:"description"`
	PublicKey   string `json:"public_key"`
}
