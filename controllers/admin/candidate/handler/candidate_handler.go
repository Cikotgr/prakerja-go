package handler

import (
	"github.com/ardin2001/backend-pemilu/controllers/admin/candidate"
	"github.com/ardin2001/backend-pemilu/controllers/admin/candidate/usecase"
	"github.com/ardin2001/backend-pemilu/helper"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type CandidateHandlerInterface interface {
	GetAll(c echo.Context) error
	GetById(c echo.Context) error
	Create(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
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
		return c.JSON(statusCode, helper.ResponseData("failed to get all candidates data", statusCode, nil))
	}
	return c.JSON(statusCode, helper.ResponseData("successful to get all candidates data", statusCode, candidates))
}

func (ch *CandidateHandlerStruct) GetById(c echo.Context) error {
	id := c.Param("id")

	candidate, statusCode, err := ch.CandidateUsecase.GetById(id)
	if err != nil {
		return c.JSON(statusCode, helper.ResponseData("failed to get candidate detail data", statusCode, nil))
	}
	return c.JSON(statusCode, helper.ResponseData("successful to get candidate detail data", statusCode, candidate))
}

func (ch *CandidateHandlerStruct) Create(c echo.Context) error {
	var candidate candidate.CreateCandidate
	id := uuid.New().String()
	c.Bind(&candidate)
	candidate.ID = id

	err := isRequestValid(candidate)
	if err != nil {
		return c.JSON(401, helper.ResponseData(err.Error(), 401, nil))
	}

	statusCode, err := ch.CandidateUsecase.Create(&candidate)
	if err != nil {
		return c.JSON(statusCode, helper.ResponseData("failed to create candidate data", statusCode, nil))
	}
	return c.JSON(statusCode, helper.ResponseData("successful to create candidate data", statusCode, nil))
}

func (ch *CandidateHandlerStruct) Update(c echo.Context) error {
	id := c.Param("id")
	var candidate candidate.UpdateCandidate
	c.Bind(&candidate)
	candidate.ID = id

	statusCode, err := ch.CandidateUsecase.Update(&candidate)
	if err != nil {
		return c.JSON(statusCode, helper.ResponseData("failed to update candidate data", statusCode, nil))
	}
	return c.JSON(statusCode, helper.ResponseData("successful to update candidate data", statusCode, nil))
}

func (ch *CandidateHandlerStruct) Delete(c echo.Context) error {
	id := c.Param("id")

	statusCode, err := ch.CandidateUsecase.Delete(id)
	if err != nil {
		return c.JSON(statusCode, helper.ResponseData("failed to delete candidate data", statusCode, nil))
	}
	return c.JSON(statusCode, helper.ResponseData("successful to delete candidate data", statusCode, nil))
}
