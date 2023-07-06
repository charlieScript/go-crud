package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// gorm.Model definition
type Post struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Title     string
	Body      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
