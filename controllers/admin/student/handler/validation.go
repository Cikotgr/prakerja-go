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
				} else if err.Field() == "NIM" {
					return errors.New("nim is required")
				}
			} else if err.Tag() == "min" {
				return errors.New("min nim should be 202410102000")
			} else if err.Tag() == "max" {
				return errors.New("max nim should be 232410102100")
			} else if err.Tag() == "number" {
				return errors.New("role id must be number")
			}
		}
	}

	return nil
}
