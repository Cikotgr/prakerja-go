package student

import "time"

type StudentResponse struct {
	ID        string    `gorm:"primarykey" json:"id"`
	NIM       string    `json:"nim"`
	RoleId    int       `json:"role_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type StudentJWTResponse struct {
	ID     string `gorm:"primarykey" json:"id"`
	NIM    string `json:"nim"`
	RoleId int    `json:"role_id"`
	Token  string `json:"token"`
}
