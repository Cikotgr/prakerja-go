package routes

import (
	"net/http"
	"os"

	config "github.com/ardin2001/backend-pemilu/configs"
	admin "github.com/ardin2001/backend-pemilu/controllers/admin/admin"
	AdminHandler "github.com/ardin2001/backend-pemilu/controllers/admin/admin/handler"
	AdminRepository "github.com/ardin2001/backend-pemilu/controllers/admin/admin/repository"
	AdminUsecase "github.com/ardin2001/backend-pemilu/controllers/admin/admin/usecase"
	AdminCandidateHandler "github.com/ardin2001/backend-pemilu/controllers/admin/candidate/handler"
	AdminCandidateRepository "github.com/ardin2001/backend-pemilu/controllers/admin/candidate/repository"
	AdminCandidateUsecase "github.com/ardin2001/backend-pemilu/controllers/admin/candidate/usecase"
	AdminStudentHandler "github.com/ardin2001/backend-pemilu/controllers/admin/student/handler"
	AdminStudentRepository "github.com/ardin2001/backend-pemilu/controllers/admin/student/repository"
	AdminStudentUsecase "github.com/ardin2001/backend-pemilu/controllers/admin/student/usecase"
	StudentCandidateHandler "github.com/ardin2001/backend-pemilu/controllers/student/candidate/handler"
	StudentCandidateRepository "github.com/ardin2001/backend-pemilu/controllers/student/candidate/repository"
	StudentCandidateUsecase "github.com/ardin2001/backend-pemilu/controllers/student/candidate/usecase"
	"github.com/ardin2001/backend-pemilu/controllers/student/student"
	StudentHandler "github.com/ardin2001/backend-pemilu/controllers/student/student/handler"
	StudentRepository "github.com/ardin2001/backend-pemilu/controllers/student/student/repository"
	StudentUsecase "github.com/ardin2001/backend-pemilu/controllers/student/student/usecase"
	"github.com/ardin2001/backend-pemilu/helper"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func StartApp() *echo.Echo {

	// new object controllers
	DB, _ := config.ConfigDatabase()

	AdminStudentR := AdminStudentRepository.NewStudentRepository(DB)
	AdminStudentU := AdminStudentUsecase.NewStudentUsecase(AdminStudentR)
	AdminStudentH := AdminStudentHandler.NewStudentHandler(AdminStudentU)
	AdminR := AdminRepository.NewAdminRepository(DB)
	AdminU := AdminUsecase.NewAdminUsecase(AdminR)
	AdminH := AdminHandler.NewAdminHandler(AdminU)
	AdminCandidateR := AdminCandidateRepository.NewCandidateRepository(DB)
	AdminCandidateU := AdminCandidateUsecase.NewCandidateUsecase(AdminCandidateR)
	AdminCandidateH := AdminCandidateHandler.NewCandidateHandler(AdminCandidateU)

	StudentR := StudentRepository.NewStudentRepository(DB)
	StudentU := StudentUsecase.NewStudentUsecase(StudentR)
	StudentH := StudentHandler.NewStudentHandler(StudentU)
	StudentCandidateR := StudentCandidateRepository.NewCandidateRepository(DB)
	StudentCandidateU := StudentCandidateUsecase.NewCandidateUsecase(StudentCandidateR)
	StudentCandidateH := StudentCandidateHandler.NewCandidateHandler(StudentCandidateU)

	e := echo.New()
	godotenv.Load()
	secretKeyAdmin := os.Getenv("SECRET_KEY_ADMIN")
	configAdmin := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(admin.JwtCustomClaimsAdmin)
		},
		SigningKey: []byte(secretKeyAdmin),
		ErrorHandler: func(c echo.Context, err error) error {
			return c.JSON(401, helper.ResponseData("failed or expired jwt token admin", 401, nil))
		},
	}

	secretKeyStudent := os.Getenv("SECRET_KEY_STUDENT")
	configStudent := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(student.JwtCustomClaimsStudent)
		},
		SigningKey: []byte(secretKeyStudent),
		ErrorHandler: func(c echo.Context, err error) error {
			return c.JSON(401, helper.ResponseData("failed or expired jwt token student", 401, nil))
		},
	}

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	routeAdmin := e.Group("/admins")
	routeAdmin.POST("/login", AdminH.LoginAdmin)
	routeAdmin.Use(echojwt.WithConfig(configAdmin))
	routeAdmin.GET("/details", AdminH.GetById)
	routeAdmin.GET("/students", AdminStudentH.GetAll)
	routeAdmin.GET("/students/:id", AdminStudentH.GetById)
	routeAdmin.POST("/students", AdminStudentH.Create)
	routeAdmin.PUT("/students/:id", AdminStudentH.Update)
	routeAdmin.DELETE("/students/:id", AdminStudentH.Delete)
	routeAdmin.GET("/candidates", AdminCandidateH.GetAll)
	routeAdmin.GET("/candidates/:id", AdminCandidateH.GetById)
	routeAdmin.POST("/candidates", AdminCandidateH.Create)
	routeAdmin.PUT("/candidates/:id", AdminCandidateH.Update)
	routeAdmin.DELETE("/candidates/:id", AdminCandidateH.Delete)

	routeStudent := e.Group("/students")
	routeStudent.POST("/login", StudentH.LoginStudent)
	routeStudent.Use(echojwt.WithConfig(configStudent))
	routeStudent.GET("/details", StudentH.GetById)
	routeStudent.GET("/candidates", StudentCandidateH.GetAll)
	routeStudent.GET("/candidates/:id", StudentCandidateH.GetById)

	return e
}
