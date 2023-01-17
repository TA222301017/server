package models

type Personel struct {
	BaseModel
	IDNumber    string `json:"id_number" gorm:"unique"`
	Name        string `json:"name"`
	Status      bool   `json:"status"`
	RoleID      uint64 `json:"-"`
	Role        Role   `json:"role,omitempty"`
	Description string `json:"description"`
	KeyID       uint64 `json:"-"`
	Key         Key    `json:"key,omitempty"`
}
