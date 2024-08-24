package models

import (
	"time"
)

type Post struct {
	ID         uint    `gorm:"primaryKey"`
	UserID     uint    `json:"user_id"`
	CategoryID uint    `json:"category_id"`
	Title      *string `json:"title" validate:"required"`
	Content    *string `json:"content" validate:"required"`
	CreatedAt  *time.Time
	UpdatedAt  *time.Time

	User     User     `json:"user" gorm:"foreignKey:UserID;references:ID"`
	Category Category `json:"category" gorm:"foreignKey:CategoryID;references:ID"`
}
