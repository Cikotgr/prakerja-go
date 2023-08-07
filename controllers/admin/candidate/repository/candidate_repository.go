package repository

import (
	"github.com/ardin2001/backend-pemilu/controllers/admin/candidate"
	"github.com/ardin2001/backend-pemilu/models"
	"gorm.io/gorm"
)

type CandidateRepositoryInterface interface {
	GetAll() (*[]candidate.CandidateResponse, error)
	GetById(id string) (*candidate.CandidateResponse, error)
	Create(candidate *models.Candidate) error
	Update(candidate *models.Candidate) error
	Delete(id string) error
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

func (cr *CandidateRepositoryStruct) Create(candidate *models.Candidate) error {
	err := cr.DB.Create(candidate).Error
	if err != nil {
		return err
	}
	return nil
}

func (cr *CandidateRepositoryStruct) Update(candidate *models.Candidate) error {
	err := cr.DB.Where("id = ?", candidate.ID).Updates(candidate).Error
	if err != nil {
		return err
	}
	return nil
}

func (cr *CandidateRepositoryStruct) Delete(id string) error {
	err := cr.DB.Transaction(func(tx *gorm.DB) error {
		err := tx.Model(&models.CandidateStudent{}).Unscoped().Delete(&models.CandidateStudent{}, "candidate_id = ?", id).Error

		if err != nil {
			return err
		}
		err = tx.Model(&models.Candidate{}).Unscoped().Delete(&models.Candidate{}, "id = ?", id).Error

		if err != nil {
			return err
		}

		return nil
	})

	return err
}
