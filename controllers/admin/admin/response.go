package admin

import "time"

type AdminResponse struct {
	ID        string    `gorm:"primarykey" json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	RoleId    int       `json:"role_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type AdminDetailResponse struct {
	ID       string `gorm:"primarykey" json:"id"`
	Username string `json:"username"`
	RoleId   int    `json:"role_id"`
}

type AdminJWTResponse struct {
	ID       string `gorm:"primarykey" json:"id"`
	Username string `json:"username"`
	RoleId   int    `json:"role_id"`
	Token    string `json:"token"`
}
