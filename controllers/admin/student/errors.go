package student

import (
	"errors"
)

var (
	// ErrFailedGetReadingList       = errors.New("failed to get all reading list data")
	// ErrFailedGetDetailReadingList = errors.New("failed to get reading list data details")
	// ErrFailedCreateReadingList    = errors.New("failed created reading list data")
	// ErrFailedUpdateReadingList    = errors.New("failed to updated reading list data")
	// ErrFailedDeleteReadingList    = errors.New("failed to delete reading list data")
	// ErrPageNotFound               = errors.New("page not found")

	ErrRequiredId  = errors.New("id is required")
	ErrRequiredNIM = errors.New("nim is required")
	ErrRequired    = errors.New("all fields are required")
	ErrMinNIM      = errors.New("min nim should be 202410102000")
	ErrMaxNIM      = errors.New("max nim should be 232410102100")
	ErrRoleId      = errors.New("role id must be number")
)
