package models

type Lock struct {
	BaseModel
	LockID      string  `json:"lock_id" gorm:"unique"`
	IpAddress   string  `json:"ip_address"`
	Label       string  `json:"name"`
	Location    string  `json:"location"`
	Description string  `json:"description"`
	Status      bool    `json:"status"`
	PublicKey   string  `json:"public_key,omitempty"`
	Plan        Plan    `json:"plan,omitempty"`
	PlanID      uint64  `json:"map_id"`
	CoordX      float32 `json:"coord_x"`
	CoordY      float32 `json:"coord_y"`
}
