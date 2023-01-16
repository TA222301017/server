package models

import "time"

type AccessRule struct {
	BaseModel
	PersonelID uint64    `json:"-"`
	Personel   Personel  `json:"personel"`
	LockID     uint64    `json:"-"`
	Lock       Lock      `json:"lock"`
	KeyID      uint64    `json:"-"`
	Key        Key       `json:"key"`
	CreatorID  uint64    `json:"-"`
	Creator    User      `json:"creator"`
	StartsAt   time.Time `json:"starts_at"`
	EndsAt     time.Time `json:"ends_at"`
}
