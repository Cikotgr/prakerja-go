package repository

import (
	"github.com/ardin2001/backend-pemilu/controllers/admin/vote"
	"gorm.io/gorm"
)

type VoteRepositoryInterface interface {
	GetAll() (*[]vote.VoteResponse, error)
}

type VoteRepositoryStruct struct {
	DB *gorm.DB
}

func NewVoteRepository(DB *gorm.DB) VoteRepositoryInterface {
	return &VoteRepositoryStruct{
		DB: DB,
	}
}

func (vr *VoteRepositoryStruct) GetAll() (*[]vote.VoteResponse, error) {
	var vote []vote.VoteResponse
	err := vr.DB.Table("candidate_students").Select("candidates.id,candidates.full_name,COUNT(candidate_students.id) as total_votes").Joins("right join candidates ON candidate_students.candidate_id = candidates.id").Group("candidate_students.candidate_id").Find(&vote).Error
	if err != nil {
		return nil, err
	}
	return &vote, nil
}
