package handler

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

func isRequestValid(m interface{}) error {

	validate := validator.New()
	err := validate.Struct(m)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println("err :", err.Tag())
			if err.Tag() == "required" {
				if err.Field() == "ID" {
					return errors.New("id is required")
				} else if err.Field() == "FullName" {
					return errors.New("fullname is required")
				} else if err.Field() == "Image" {
					return errors.New("image is required")
				} else if err.Field() == "Batch" {
					return errors.New("batch is required")
				} else if err.Field() == "Vision" {
					return errors.New("vision is required")
				} else if err.Field() == "Mission" {
					return errors.New("mission is required")
				}
			}
		}
	}

	return nil
}
