package template

type PersonelData struct {
	ID          uint64 `json:"id"`
	Name        string `json:"name"`
	PersonelID  string `json:"personel_id"`
	Status      bool   `json:"status"`
	Role        string `json:"role"`
	Description string `json:"description"`
}

type AddPersonelRequest struct {
	Name        string `json:"name"`
	PersonelID  string `json:"personel_id"`
	RoleID      uint64 `json:"role_id"`
	KeyID       uint64 `json:"key_id"`
	Status      bool   `json:"status"`
	Description string `json:"description"`
}

func (a AddPersonelRequest) Validate() error {
	return nil
}

type EditPersonelRequest struct {
	Name        string `json:"name"`
	PersonelID  string `json:"personel_id"`
	RoleID      uint64 `json:"role_id"`
	KeyID       uint64 `json:"key_id"`
	Status      bool   `json:"status"`
	Description string `json:"description"`
}
