package models

import "time"

type HealthcheckLog struct {
	BaseModel
	LockID    uint64    `json:"lock_id"`
	Lock      Lock      `json:"-"`
	Timestamp time.Time `json:"timestamp"`
	Status    bool      `json:"status"`
}
