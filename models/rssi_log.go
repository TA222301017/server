package models

import "time"

type RSSILog struct {
	BaseModel
	RSSI      int       `json:"rssi"`
	LockID    uint64    `json:"-"`
	Lock      Lock      `json:"lock"`
	KeyID     uint64    `json:"-"`
	Key       Key       `json:"key"`
	Timestamp time.Time `json:"timestamp"`
}
