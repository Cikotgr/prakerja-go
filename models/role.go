package models

import "time"

type Role struct {
	ID          uint `gorm:"primarykey"`
	Description string
	Admins      []Admin   `gorm:"foreignKey:RoleId"`
	Students    []Student `gorm:"foreignKey:RoleId"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
