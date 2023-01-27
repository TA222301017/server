package models

import "time"

type RSSILog struct {
	BaseModel
	RSSI       int       `json:"rssi"`
	PersonelID uint64    `json:"-"`
	Personel   Personel  `json:"personel"`
	LockID     uint64    `json:"-"`
	Lock       Lock      `json:"lock"`
	KeyID      uint64    `json:"-"`
	Key        Key       `json:"key"`
	Timestamp  time.Time `json:"timestamp"`
}
