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

	return nil
}

type EditAccessRule struct {
	LockID   uint64    `json:"lock_id"`
	StartsAt time.Time `json:"starts_at"`
	EndsAt   time.Time `json:"ends_at"`
}

type AccessRuleData struct {
	ID       uint64    `json:"id"`
	StartsAt time.Time `json:"starts_at"`
	EndsAt   time.Time `json:"ends_at"`
	Lock     string    `json:"lock"`
}
