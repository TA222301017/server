package template

import "errors"

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (l LoginRequest) Validate() error {
	if l.Username == "" {
		return errors.New("username must not be empty")
	}

	if l.Password == "" {
		return errors.New("password must not be empty")
	}

	return nil
}

type RegisterRequest struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (r RegisterRequest) Validate() error {
	if r.Name == "" {
		return errors.New("name must not be empty")
	}

	if r.Username == "" {
		return errors.New("username must not be empty")
	}

	if r.Password == "" {
		return errors.New("password must not be empty")
	}

	return nil
}
