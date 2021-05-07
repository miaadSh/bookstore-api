package models

import (
	"time"
)

type (
	Base struct {
		ID        uint `gorm:"primarykey"`
		CreatedAt time.Time
		UpdateAt  time.Time
		DeletedAt *time.Time `gorm:"index"`
	}
)
