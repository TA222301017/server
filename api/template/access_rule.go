package template

import (
	"errors"
	"time"
)

type AddAccessRule struct {
	LockID     uint64    `json:"lock_id"`
	PersonelID uint64    `json:"personel_id"`
	StartsAt   time.Time `json:"starts_at"`
	EndsAt     time.Time `json:"ends_at"`
}

func (a AddAccessRule) Validate() error {
	if a.LockID == 0 {
		return errors.New("lock_id must not be emtpy")
	}

	if a.PersonelID == 0 {
		return errors.New("personel_id must not be empty")
	}

	if a.StartsAt.Unix() == 0 {
		return errors.New("starts_at must not be empty")
	}

	if a.EndsAt.Unix() == 0 {
		return errors.New("ends_at must not be empty")
	}

	if a.EndsAt.Unix() < a.StartsAt.Unix() {
		return errors.New("ends_at must be after starts_at")
	}

	return nil
}

type EditAccessRule struct {
	LockID   uint64    `json:"lock_id"`
	StartsAt time.Time `json:"starts_at"`
	EndsAt   time.Time `json:"ends_at"`
}

func (e EditAccessRule) Validate() error {
	if e.LockID == 0 {
		return errors.New("lock_id must not be emtpy")
	}

	if e.StartsAt.Unix() == 0 {
		return errors.New("starts_at must not be empty")
	}

	if e.EndsAt.Unix() == 0 {
		return errors.New("ends_at must not be empty")
	}

	if e.EndsAt.Unix() < e.StartsAt.Unix() {
		return errors.New("ends_at must be after starts_at")
	}

	return nil
}

type AccessRuleData struct {
	ID         uint64    `json:"id"`
	StartsAt   time.Time `json:"starts_at"`
	EndsAt     time.Time `json:"ends_at"`
	LockID     uint64    `json:"lock_id,omitempty"`
	Lock       string    `json:"lock"`
	Location   string    `json:"location"`
	PersonelID uint64    `json:"personel_id,omitempty"`
	Personel   string    `json:"personel,omitempty"`
	KeyID      uint64    `json:"key_id,omitempty"`
	Key        string    `json:"key,omitempty"`
}
