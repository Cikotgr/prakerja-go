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

	ErrRequiredUsername = errors.New("username is required")
	ErrRequiredPassword = errors.New("password is required")
	ErrRequired         = errors.New("all fields are required")
	ErrMinUsername      = errors.New("username min 6")
	ErrMinPassword      = errors.New("password min 6")
	ErrMin              = errors.New("all field min 6")
)
