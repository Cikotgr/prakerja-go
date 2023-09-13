package handler

import (
	"github.com/ardin2001/backend-pemilu/controllers/student/student"
	"github.com/ardin2001/backend-pemilu/controllers/student/vote"
	"github.com/ardin2001/backend-pemilu/controllers/student/vote/usecase"
	"github.com/ardin2001/backend-pemilu/helper"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type VoteHandlerInterface interface {
	Create(c echo.Context) error
}

type VoteHandlerStruct struct {
	VoteUsecase usecase.VoteUsecaseInterface
}

func NewVoteHandler(VoteUsecase usecase.VoteUsecaseInterface) VoteHandlerInterface {
	return &VoteHandlerStruct{
		VoteUsecase: VoteUsecase,
	}
}

func (vh *VoteHandlerStruct) Create(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*student.JwtCustomClaimsStudent)
	var vote vote.CreateVote
	id := uuid.New().String()
	c.Bind(&vote)
	vote.ID = id
	vote.StudentId = claims.ID

	err := isRequestValid(vote)
	if err != nil {
		return c.JSON(401, helper.ResponseData(err.Error(), 401, nil))
	}

	statusCode, err := vh.VoteUsecase.Create(&vote)
	if err != nil {
		return c.JSON(statusCode, helper.ResponseData(err.Error(), statusCode, nil))
	}
	return c.JSON(statusCode, helper.ResponseData("successful to create vote data", statusCode, nil))
}
