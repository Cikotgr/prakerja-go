package handler

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

func isRequestValid(m interface{}) error {

	validate := validator.New()
	err := validate.Struct(m)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			if err.Tag() == "required" {
				if err.Field() == "ID" {
					return errors.New("id is required")
				} else if err.Field() == "CandidateId" {
					return errors.New("candidate id is required")
				} else if err.Field() == "StudentId" {
					return errors.New("student id is required")
				}
			}
		}
	}

	return nil
}
