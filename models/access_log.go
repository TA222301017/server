package models

import "time"

type AccessLog struct {
	BaseModel
	PersonelName     string    `json:"personel_name"`
	PersonelIDNumber string    `json:"personel_id_number"`
	LockID           uint64    `json:"-"`
	Lock             Lock      `json:"lock"`
	Location         string    `json:"location"`
	KeyID            uint64    `json:"-"`
	Key              Key       `json:"key"`
	Timestamp        time.Time `json:"timestamp"`
}
