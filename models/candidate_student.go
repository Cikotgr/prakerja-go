package models

import (
	"time"
)

type CandidateStudent struct {
	ID          string `gorm:"primarykey"`
	CandidateId string `gorm:"type:varchar(50);index"`
	StudentId   string `gorm:"type:varchar(50);index:,unique"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
