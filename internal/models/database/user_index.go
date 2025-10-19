package database

import (
	"time"
)

type UserIndex struct {
	ID             uint      `gorm:"primaryKey"`
	UserId         string    `gorm:"size:100;not null"`
	IndexName      string    `gorm:"size:100;not null"`
	IndexStructure JSONB     `gorm:"type:jsonb"`
	CreatedAt      time.Time `gorm:"autoCreateTime"`
}
