package usecase

import (
	"net/http"

	"github.com/ardin2001/backend-pemilu/controllers/student/vote"
	"github.com/ardin2001/backend-pemilu/controllers/student/vote/repository"
	"github.com/ardin2001/backend-pemilu/models"
	"github.com/pkg/errors"
)

type VoteUsecaseInterface interface {
	Create(vote *vote.CreateVote) (int, error)
}

type VoteUsecaseStruct struct {
	VoteRepository repository.VoteRepositoryInterface
}

func NewVoteUsecase(VoteRepository repository.VoteRepositoryInterface) VoteUsecaseInterface {
	return &VoteUsecaseStruct{
		VoteRepository: VoteRepository,
	}
}

func (vu *VoteUsecaseStruct) Create(vote *vote.CreateVote) (int, error) {
	newStudent := models.CandidateStudent{
		ID:          vote.ID,
		CandidateId: vote.CandidateId,
		StudentId:   vote.StudentId,
	}

	err := vu.VoteRepository.Create(&newStudent)
	if err != nil {
		return http.StatusBadRequest, errors.New("failed to create vote data")
	}
	return http.StatusOK, nil
}
