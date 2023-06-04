package template

import (
	"encoding/hex"
	"errors"
)

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
	AESKey      string `json:"aes_key"`
}

func (a AddKeyRequest) Validate() error {
	if a.Name == "" {
		return errors.New("name must not be empty")
	}

	if a.KeyID == "" {
		return errors.New("key_id must not be empty")
	}

	if len(a.KeyID) > 16 {
		return errors.New("key_id must have length no more than 16 characters")
	}

	if len(a.AESKey) != 32 {
		return errors.New("aes_key must have length of 32 characters")
	}

	if _, err := hex.DecodeString(a.AESKey); err != nil {
		return errors.New("aes_key is not valid hexadecimal")
	}

	return nil
}

type EditKeyRequest struct {
	Name        string `json:"name"`
	KeyID       string `json:"key_id"`
	Status      bool   `json:"status"`
	Description string `json:"description"`
	AESKey      string `json:"aes_key"`
}

func (e EditKeyRequest) Validate() error {
	if len(e.KeyID) != 0 {
		if len(e.KeyID) > 16 {
			return errors.New("key_id must have length no more than 16 characters")
		}
	}

	if len(e.AESKey) != 0 {
		if len(e.AESKey) != 32 {
			return errors.New("aes_key must have length of 32 characters")
		}

		if _, err := hex.DecodeString(e.AESKey); err != nil {
			return errors.New("aes_key is not valid hexadecimal")
		}
	}

	return nil
}
