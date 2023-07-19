package models

import "time"

type Admin struct {
	ID        string `gorm:"primarykey"`
	Username  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
