package handler

import (
	"github.com/ardin2001/backend-pemilu/controllers/admin/vote/usecase"
	"github.com/ardin2001/backend-pemilu/helper"
	"github.com/labstack/echo/v4"
)

type VoteHandlerInterface interface {
	GetAll(c echo.Context) error
}

type VoteHandlerStruct struct {
	VoteUsecase usecase.VoteUsecaseInterface
}

func NewVoteHandler(VoteUsecase usecase.VoteUsecaseInterface) VoteHandlerInterface {
	return &VoteHandlerStruct{
		VoteUsecase: VoteUsecase,
	}
}

func (vh *VoteHandlerStruct) GetAll(c echo.Context) error {
	votes, statusCode, err := vh.VoteUsecase.GetAll()
	if err != nil {
		return c.JSON(statusCode, helper.ResponseData(err.Error(), statusCode, nil))
	}
	return c.JSON(statusCode, helper.ResponseData("successful to create vote data", statusCode, votes))
}
