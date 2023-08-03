package models

import (
	"time"
)

type Admin struct {
	ID        string `gorm:"primarykey"`
	Username  string
	Password  string
	RoleId    int `gorm:"type:int(50);index"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
