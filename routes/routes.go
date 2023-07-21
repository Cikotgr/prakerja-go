package routes

import (
	"net/http"

	config "github.com/ardin2001/backend-pemilu/configs"
	AdminStudentHandler "github.com/ardin2001/backend-pemilu/controllers/admin/student/handler"
	AdminStudentRepository "github.com/ardin2001/backend-pemilu/controllers/admin/student/repository"
	AdminStudentUsecase "github.com/ardin2001/backend-pemilu/controllers/admin/student/usecase"
	"github.com/labstack/echo/v4"
)

func StartApp() *echo.Echo {

	// new object controllers
	DB, _ := config.ConfigDatabase()

	AdminStudentR := AdminStudentRepository.NewStudentRepository(DB)
	AdminStudentU := AdminStudentUsecase.NewStudentUsecase(AdminStudentR)
	AdminStudentH := AdminStudentHandler.NewStudentHandler(AdminStudentU)

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/students", AdminStudentH.GetAll)
	e.GET("/students/:id", AdminStudentH.GetById)
	e.POST("/students", AdminStudentH.Create)
	e.PUT("/students/:id", AdminStudentH.Update)
	e.DELETE("/students/:id", AdminStudentH.Delete)

	return e
}
