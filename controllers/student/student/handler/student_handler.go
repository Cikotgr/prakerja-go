package handler

import (
	"github.com/ardin2001/backend-pemilu/controllers/student/student"
	"github.com/ardin2001/backend-pemilu/controllers/student/student/usecase"
	"github.com/ardin2001/backend-pemilu/helper"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type StudentHandlerInterface interface {
	GetById(c echo.Context) error
	LoginStudent(c echo.Context) error
}

type StudentHandlerStruct struct {
	StudentUsecase usecase.StudentUsecaseInterface
}

func NewStudentHandler(StudentUsecase usecase.StudentUsecaseInterface) StudentHandlerInterface {
	return &StudentHandlerStruct{
		StudentUsecase: StudentUsecase,
	}
}

func (ah *StudentHandlerStruct) GetById(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*student.JwtCustomClaimsStudent)
	studentJWTResponse, statusCode, err := ah.StudentUsecase.LoginStudent(claims.NIM)
	if err != nil {
		return c.JSON(statusCode, helper.ResponseData("failed to login student", statusCode, nil))
	}
	return c.JSON(statusCode, helper.ResponseData("successful to login student", statusCode, studentJWTResponse))
}

func (ah *StudentHandlerStruct) LoginStudent(c echo.Context) error {
	var student student.RequestLoginStudent
	c.Bind(&student)

	err := isRequestValid(student)
	if err != nil {
		return c.JSON(401, helper.ResponseData(err.Error(), 401, nil))
	}

	studentJWTResponse, statusCode, err := ah.StudentUsecase.LoginStudent(student.NIM)
	if err != nil {
		return c.JSON(statusCode, helper.ResponseData("failed to login student", statusCode, nil))
	}
	return c.JSON(statusCode, helper.ResponseData("successful to login student", statusCode, studentJWTResponse))
}
