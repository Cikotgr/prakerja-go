package repository

import (
	"github.com/ardin2001/backend-pemilu/models"
	"gorm.io/gorm"
)

type VoteRepositoryInterface interface {
	Create(vote *models.CandidateStudent) error
}

type VoteRepositoryStruct struct {
	DB *gorm.DB
}

func NewVoteRepository(DB *gorm.DB) VoteRepositoryInterface {
	return &VoteRepositoryStruct{
		DB: DB,
	}
}

func (vr *VoteRepositoryStruct) Create(vote *models.CandidateStudent) error {
	err := vr.DB.Create(vote).Error
	if err != nil {
		return err
	}
	return nil
}
