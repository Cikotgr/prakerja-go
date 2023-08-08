package handler

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

func isRequestValid(m interface{}) error {

	validate := validator.New()
	err := validate.Struct(m)

	fmt.Println(m, err)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			if err.Tag() == "required" {
				if err.Field() == "Username" && err.Field() == "Password" {
					return errors.New("all fields are required")
				} else if err.Field() == "Username" {
					return errors.New("username is required")
				} else if err.Field() == "Password" {
					return errors.New("password is required")
				}
			} else if err.Tag() == "min" {
				if err.Field() == "Username" && err.Field() == "Password" {
					return errors.New("all field min 6")
				} else if err.Field() == "Username" {
					return errors.New("username min 6")
				} else if err.Field() == "Password" {
					return errors.New("password min 6")
				}
			}
		}
	}

	return nil
}
