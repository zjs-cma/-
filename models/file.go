package models

import (
	"time"

	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(&File{})
}

type File struct {
	ID        string    `gorm:"primaryKey;size:36"`
	Name      string    `gorm:"size:255;not null"`
	Path      string    `gorm:"size:512;not null"`
	Size      int64     `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}
