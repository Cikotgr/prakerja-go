package handler

import (
	"fmt"

	"github.com/ardin2001/backend-pemilu/controllers/admin/student"
	"github.com/go-playground/validator/v10"
)

func isRequestValid(m interface{}) error {

	validate := validator.New()
	err := validate.Struct(m)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println("err :", err.Tag())
			if err.Tag() == "required" {
				if err.Field() == "ID" && err.Field() == "NIM" {
					return student.ErrRequired
				} else if err.Field() == "ID" {
					return student.ErrRequiredId
				} else if err.Field() == "NIM" {
					return student.ErrRequiredNIM
				}
			}
		}
	}

	return nil
}
