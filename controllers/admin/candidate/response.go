package candidate

import "time"

type CandidateResponse struct {
	ID        string    `gorm:"primarykey"`
	FullName  string    `json:"fullname"`
	Image     string    `json:"image"`
	Batch     int       `json:"batch"`
	Vision    string    `json:"vision"`
	Mission   string    `json:"mission"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
