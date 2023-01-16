package models

import "time"

type BaseModel struct {
	ID        uint64    `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_at,omitempty" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at,omitempty" gorm:"autoUpdateTime"`
}
