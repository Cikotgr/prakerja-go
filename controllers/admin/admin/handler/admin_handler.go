package handler

import (
	"github.com/ardin2001/backend-pemilu/controllers/admin/admin"
	"github.com/ardin2001/backend-pemilu/controllers/admin/admin/usecase"
	"github.com/ardin2001/backend-pemilu/helper"
	"github.com/labstack/echo/v4"
)

type AdminHandlerInterface interface {
	LoginAdmin(c echo.Context) error
}

type AdminHandlerStruct struct {
	AdminUsecase usecase.AdminUsecaseInterface
}

func NewAdminHandler(AdminUsecase usecase.AdminUsecaseInterface) AdminHandlerInterface {
	return &AdminHandlerStruct{
		AdminUsecase: AdminUsecase,
	}
}

func (ah *AdminHandlerStruct) LoginAdmin(c echo.Context) error {
	var admin admin.RequestLoginAdmin
	c.Bind(&admin)

	err := isRequestValid(admin)
	if err != nil {
		return c.JSON(401, helper.ResponseData(err.Error(), 401, nil))
	}

	adminJWTResponse, statusCode, err := ah.AdminUsecase.LoginAdmin(admin.Username, admin.Password)
	if err != nil {
		return c.JSON(statusCode, helper.ResponseData(err.Error(), statusCode, nil))
	}
	return c.JSON(statusCode, helper.ResponseData("successful to login", statusCode, adminJWTResponse))
}
