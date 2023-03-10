package template

import "time"

type RSSILogData struct {
	Timestamp  time.Time `json:"timestamp"`
	PersonelID uint64    `json:"personel_id"`
	Personel   string    `json:"personel"`
	LockID     uint64    `json:"lock_id"`
	Lock       string    `json:"lock"`
	KeyID      uint64    `json:"key_id"`
	Key        string    `json:"key"`
	Location   string    `json:"location"`
	RSSI       int       `json:"rssi"`
}

type HealthcheckLogData struct {
	ID        uint64    `json:"id"`
	Device    string    `json:"device"`
	DeviceID  uint64    `json:"device_id"`
	Location  string    `json:"location"`
	Status    bool      `json:"status"`
	Timestamp time.Time `json:"timestamp"`
}

type AccessLogData struct {
	ID         uint64    `json:"id"`
	Personel   string    `json:"personel"`
	PersonelID uint64    `json:"personel_id"`
	Lock       string    `json:"lock"`
	LockID     uint64    `json:"lock_id"`
	Key        string    `json:"key"`
	KeyID      uint64    `json:"key_id"`
	Location   string    `json:"location"`
	Timestamp  time.Time `json:"timestamp"`
}
