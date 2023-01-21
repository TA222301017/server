package template

import "errors"

type KeyData struct {
	ID      uint64 `json:"id"`
	KeyID   string `json:"key_id"`
	Name    string `json:"name"`
	Status  bool   `json:"status"`
	Owner   string `json:"owner"`
	OwnerID uint64 `json:"owner_id"`
}

type AddKeyRequest struct {
	Name        string `json:"name"`
	KeyID       string `json:"key_id"`
	Status      bool   `json:"status"`
	Description string `json:"description"`
}

func (a AddKeyRequest) Validate() error {
	if a.Name == "" {
		return errors.New("name must not be empty")
	}

	if a.KeyID == "" {
		return errors.New("key_id must not be empty")
	}

	if len(a.KeyID) != 32 {
		return errors.New("key_id must have length of 32 characters")
	}

	return nil
}

type EditKeyRequest struct {
	Name        string `json:"name"`
	KeyID       string `json:"key_id"`
	Status      bool   `json:"status"`
	Description string `json:"description"`
}

func (e EditKeyRequest) Validate() error {
	if len(e.KeyID) != 32 {
		return errors.New("key_id must have length of 32 characters")
	}

	return nil
}
