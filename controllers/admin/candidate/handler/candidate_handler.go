package handler

import "github.com/labstack/echo/v4"

type CandidateHandlerInterface interface {
	GetAll(c echo.Context) error
	GetById(c echo.Context) error
	Create(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
}

type CandidateHandlerStruct struct {
}

// func NewCandidateHandler()
