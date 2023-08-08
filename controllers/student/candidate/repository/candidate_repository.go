package repository

import (
	"github.com/ardin2001/backend-pemilu/controllers/student/candidate"
	"gorm.io/gorm"
)

type CandidateRepositoryInterface interface {
	GetAll() (*[]candidate.CandidateResponse, error)
	GetById(id string) (*candidate.CandidateResponse, error)
}

type CandidateRepositoryStruct struct {
	DB *gorm.DB
}

func NewCandidateRepository(DB *gorm.DB) CandidateRepositoryInterface {
	return &CandidateRepositoryStruct{
		DB: DB,
	}
}

func (cr *CandidateRepositoryStruct) GetAll() (*[]candidate.CandidateResponse, error) {
	var candidates []candidate.CandidateResponse

	err := cr.DB.Table("candidates").Find(&candidates).Error
	if err != nil {
		return nil, err
	}
	return &candidates, nil
}

func (cr *CandidateRepositoryStruct) GetById(id string) (*candidate.CandidateResponse, error) {
	var candidate candidate.CandidateResponse
	err := cr.DB.Table("candidates").Where("id =?", id).First(&candidate).Error
	if err != nil {
		return nil, err
	}
	return &candidate, nil
}
