package template

import "errors"

type EditAdminRequest struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type CreateAdminRequest struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (e CreateAdminRequest) Validate() error {
	if e.Name == "" {
		return errors.New("name must not be empty")
	}

	if e.Username == "" {
		return errors.New("username must not be empty")
	}

	if len(e.Password) < 8 {
		return errors.New("password must be at least 8 characters long")
	}

	return nil
}
