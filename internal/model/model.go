package model

import (
	"time"
    "github.com/google/uuid"
    "gorm.io/gorm"
)

type Movie struct {
	gorm.Model              // Adds ID, CreatedAt, UpdatedAt, DeletedAt
	ID          uuid.UUID   `gorm:"type:uuid;default:uuid_generate_v4()"` // UUID with default value
	Title       string      `gorm:"type:varchar(255);not null"`
	ReleaseDate time.Time   `gorm:"type:date"`
	Rating      int         `gorm:"type:int"`
}