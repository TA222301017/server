package template

import "time"

type RSSILogData struct {
	Timestamp    time.Time `json:"timestamp"`
	LockName     string    `json:"lock_name"`
	LockLocation string    `json:"lock_location"`
}
