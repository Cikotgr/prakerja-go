package usecase

import (
	"errors"
	"net/http"

	"github.com/ardin2001/backend-pemilu/controllers/student/candidate"
	"github.com/ardin2001/backend-pemilu/controllers/student/candidate/repository"
)

type CandidateUsecaseInterface interface {
	GetAll() (*[]candidate.CandidateResponse, int, error)
	GetById(id string) (*candidate.CandidateResponse, int, error)
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
		return nil, http.StatusInternalServerError, errors.New("failed to get all candidate data")
	}
	return candidates, http.StatusOK, nil
}

func (cu *CandidateUsecaseStruct) GetById(id string) (*candidate.CandidateResponse, int, error) {
	candidate, err := cu.CandidateRepository.GetById(id)
	if err != nil {
		return nil, http.StatusInternalServerError, errors.New("failed to get candidate data details")
	}
	return candidate, http.StatusOK, nil
}
