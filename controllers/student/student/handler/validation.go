package handler

import (
	"fmt"

	"github.com/ardin2001/backend-pemilu/controllers/admin/admin"
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
					return admin.ErrRequired
				} else if err.Field() == "Username" {
					return admin.ErrRequiredUsername
				} else if err.Field() == "Password" {
					return admin.ErrRequiredPassword
				}
			} else if err.Tag() == "min" {
				if err.Field() == "Username" && err.Field() == "Password" {
					return admin.ErrMin
				} else if err.Field() == "Username" {
					return admin.ErrMinUsername
				} else if err.Field() == "Password" {
					return admin.ErrMinPassword
				}
			}
		}
	}

	return nil
}
