package usecase

import (
	"net/http"

	"github.com/ardin2001/backend-pemilu/controllers/admin/candidate"
	"github.com/ardin2001/backend-pemilu/controllers/admin/candidate/repository"
	"github.com/ardin2001/backend-pemilu/models"
)

type CandidateUsecaseInterface interface {
	GetAll() (*[]candidate.CandidateResponse, int, error)
	GetById(id string) (*candidate.CandidateResponse, int, error)
	Create(candidate *candidate.CreateCandidate) (int, error)
	Update(candidate *candidate.UpdateCandidate) (int, error)
	Delete(id string) (int, error)
}

type CandidateUsecaseStruct struct {
	CandidateRepository repository.CandidateRepositoryInterface
}

func NewCandidateUsecase(CandidateRepository repository.CandidateRepositoryInterface) CandidateUsecaseInterface {
	return &CandidateUsecaseStruct{
		CandidateRepository: CandidateRepository,
	}
}

func (cu *CandidateUsecaseStruct) GetAll() (*[]candidate.CandidateResponse, int, error) {
	candidates, err := cu.CandidateRepository.GetAll()
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	return candidates, http.StatusOK, nil
}

func (cu *CandidateUsecaseStruct) GetById(id string) (*candidate.CandidateResponse, int, error) {
	candidate, err := cu.CandidateRepository.GetById(id)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	return candidate, http.StatusOK, nil
}

func (cu *CandidateUsecaseStruct) Create(candidate *candidate.CreateCandidate) (int, error) {
	newCandidate := models.Candidate{
		ID:       candidate.ID,
		FullName: candidate.FullName,
		Image:    candidate.Image,
		Batch:    candidate.Batch,
		Vision:   candidate.Vision,
		Mission:  candidate.Mission,
	}

	err := cu.CandidateRepository.Create(&newCandidate)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}

func (cu *CandidateUsecaseStruct) Update(candidate *candidate.UpdateCandidate) (int, error) {
	newCandidate := models.Candidate{
		ID:       candidate.ID,
		FullName: candidate.FullName,
		Image:    candidate.Image,
		Batch:    candidate.Batch,
		Vision:   candidate.Vision,
		Mission:  candidate.Mission,
	}

	err := cu.CandidateRepository.Update(&newCandidate)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}

func (cu *CandidateUsecaseStruct) Delete(id string) (int, error) {
	err := cu.CandidateRepository.Delete(id)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}
