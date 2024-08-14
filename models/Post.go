package models

import (
	"time"
)

type Post struct {
	ID        uint    `gorm:"primaryKey"`
	UserID    uint    `json:"user_id"`
	Title     *string `json:"title" validate:"required"`
	Content   *string `json:"content" validate:"required"`
	CreatedAt *time.Time
	UpdatedAt *time.Time
}
