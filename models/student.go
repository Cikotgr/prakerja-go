package models

import (
	"time"
)

type Student struct {
	ID        string `gorm:"primarykey"`
	NIM       int    `gorm:"unique"`
	RoleId    int    `gorm:"type:int(50);index"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
