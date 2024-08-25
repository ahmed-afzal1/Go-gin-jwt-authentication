package models

import (
	"time"
)

type Category struct {
	ID        uint    `gorm:"primaryKey"`
	Title     *string `json:"title" validate:"required"`
	CreatedAt *time.Time
	UpdatedAt *time.Time

	Posts []Post `json:"posts" gorm:"foreignKey:UserID;references:ID"`
}
