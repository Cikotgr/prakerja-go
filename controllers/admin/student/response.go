package student

import "time"

type StudentResponse struct {
	ID        string    `gorm:"primarykey" json:"id"`
	NIM       int       `gorm:"unique" json:"nim"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
