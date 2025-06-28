package model

import "time"

type Role struct {
	ID        int64 `gorm:"primary_key"`
	Name      string
	Users     []User `gorm:"many2many:user_role"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
