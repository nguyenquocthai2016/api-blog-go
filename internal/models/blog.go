package models

import (
	"time"

	"gorm.io/gorm"
)

type Blog struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name" gorm:"size:255;not null"`
	Slug      string         `json:"slug" gorm:"size:255;not null;unique"`
	Avatar    string         `json:"avatar" gorm:"size:255"`
	Content   string         `json:"content" gorm:"type:text"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
type BlogList struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Slug      string    `json:"slug"`
	Avatar    string    `json:"avatar"`
	CreatedAt time.Time `json:"created_at"`
}
