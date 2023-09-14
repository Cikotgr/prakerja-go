package usecase

import (
	"net/http"

	"github.com/ardin2001/backend-pemilu/controllers/admin/vote"
	"github.com/ardin2001/backend-pemilu/controllers/admin/vote/repository"
	"github.com/pkg/errors"
)

type VoteUsecaseInterface interface {
	GetAll() (*[]vote.VoteResponse, int, error)
}

type VoteUsecaseStruct struct {
	VoteRepository repository.VoteRepositoryInterface
}

func NewVoteUsecase(VoteRepository repository.VoteRepositoryInterface) VoteUsecaseInterface {
	return &VoteUsecaseStruct{
		VoteRepository: VoteRepository,
	}
}

func (vu *VoteUsecaseStruct) GetAll() (*[]vote.VoteResponse, int, error) {
	votes, err := vu.VoteRepository.GetAll()
	if err != nil {
		return nil, http.StatusBadRequest, errors.New("failed to get vote data")
	}
	return votes, http.StatusOK, nil
}
