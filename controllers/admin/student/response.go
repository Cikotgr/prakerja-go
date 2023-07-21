package student

import "time"

type StudentResponse struct {
	ID        string `gorm:"primarykey"`
	NIM       int    `gorm:"unique"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
