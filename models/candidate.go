package models

import (
	"time"
)

type Candidate struct {
	ID                string `gorm:"primarykey"`
	FullName          string
	Image             string
	Batch             int
	Vision            string
	Mission           string
	CandidateStudents []CandidateStudent `gorm:"foreignKey:CandidateId"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
}
