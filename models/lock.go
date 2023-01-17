package models

type Lock struct {
	BaseModel
	LockID      string `json:"lock_id" gorm:"unique"`
	IpAddress   string `json:"ip_address"`
	Label       string `json:"name"`
	Location    string `json:"location"`
	Description string `json:"description"`
	PublicKey   string `json:"public_key,omitempty"`
}
