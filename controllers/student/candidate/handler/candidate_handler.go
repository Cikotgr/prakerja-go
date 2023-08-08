package handler

import (
	"github.com/ardin2001/backend-pemilu/controllers/student/candidate/usecase"
	"github.com/ardin2001/backend-pemilu/helper"
	"github.com/labstack/echo/v4"
)

type CandidateHandlerInterface interface {
	GetAll(c echo.Context) error
	GetById(c echo.Context) error
}

type CandidateHandlerStruct struct {
	CandidateUsecase usecase.CandidateUsecaseInterface
}

func NewCandidateHandler(CandidateUsecase usecase.CandidateUsecaseInterface) CandidateHandlerInterface {
	return &CandidateHandlerStruct{
		CandidateUsecase: CandidateUsecase,
	}
}

func (ch *CandidateHandlerStruct) GetAll(c echo.Context) error {
	candidates, statusCode, err := ch.CandidateUsecase.GetAll()
	if err != nil {
		return c.JSON(statusCode, helper.ResponseData(err.Error(), statusCode, nil))
	}
	return c.JSON(statusCode, helper.ResponseData("successful to get all candidates data", statusCode, candidates))
}

func (ch *CandidateHandlerStruct) GetById(c echo.Context) error {
	id := c.Param("id")

	candidate, statusCode, err := ch.CandidateUsecase.GetById(id)
	if err != nil {
		return c.JSON(statusCode, helper.ResponseData(err.Error(), statusCode, nil))
	}
	return c.JSON(statusCode, helper.ResponseData("successful to get candidate detail data", statusCode, candidate))
}
