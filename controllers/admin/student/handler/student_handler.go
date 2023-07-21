package handler

import (
	"net/http"

	"github.com/ardin2001/backend-pemilu/controllers/admin/student"
	"github.com/ardin2001/backend-pemilu/controllers/admin/student/usecase"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type StudentHandlerInterface interface {
	GetAll(c echo.Context) error
	GetById(c echo.Context) error
	Create(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
}

type StudentHandlerStruct struct {
	StudentUsecase usecase.StudentUsecaseInterface
}

func NewStudentHandler(StudentUsecase usecase.StudentUsecaseInterface) StudentHandlerInterface {
	return &StudentHandlerStruct{
		StudentUsecase: StudentUsecase,
	}
}

func (sh *StudentHandlerStruct) GetAll(c echo.Context) error {
	students, err := sh.StudentUsecase.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, students)
	}
	return c.JSON(http.StatusOK, students)
}

func (sh *StudentHandlerStruct) GetById(c echo.Context) error {
	id := c.Param("id")

	student, err := sh.StudentUsecase.GetById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, student)
	}
	return c.JSON(http.StatusOK, student)
}

func (sh *StudentHandlerStruct) Create(c echo.Context) error {
	var student student.CreateStudent
	id := uuid.New().String()
	c.Bind(&student)
	student.ID = id

	err := sh.StudentUsecase.Create(&student)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
	}
	return c.JSON(http.StatusOK, nil)
}

func (sh *StudentHandlerStruct) Update(c echo.Context) error {
	id := c.Param("id")
	var student student.UpdateStudent
	c.Bind(&student)
	student.ID = id

	err := sh.StudentUsecase.Update(&student)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
	}
	return c.JSON(http.StatusOK, nil)
}

func (sh *StudentHandlerStruct) Delete(c echo.Context) error {
	id := c.Param("id")
	err := sh.StudentUsecase.Delete(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
	}
	return c.JSON(http.StatusOK, nil)
}
