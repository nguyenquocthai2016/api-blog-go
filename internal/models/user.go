package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Username  string         `json:"username" gorm:"size:255;not null;unique" binding:"required,min=3,max=255"`
	Password  string         `json:"password,omitempty" gorm:"size:255;not null" binding:"required,min=6"`
	Email     string         `json:"email" gorm:"size:255;not null;unique" binding:"required,email"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

func (User) TableName() string {
	return "user"
}
