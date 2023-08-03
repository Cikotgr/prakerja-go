package usecase

import (
	"net/http"
	"os"
	"time"

	"github.com/ardin2001/backend-pemilu/controllers/student/student"
	"github.com/ardin2001/backend-pemilu/controllers/student/student/repository"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

type StudentUsecaseInterface interface {
	LoginStudent(nim string) (*student.StudentJWTResponse, int, error)
}

type StudentUsecaseStruct struct {
	StudentRepository repository.StudentRepositoryInterface
}

func NewStudentUsecase(StudentRepository repository.StudentRepositoryInterface) StudentUsecaseInterface {
	return &StudentUsecaseStruct{
		StudentRepository: StudentRepository,
	}
}

func (au *StudentUsecaseStruct) LoginStudent(nim string) (*student.StudentJWTResponse, int, error) {
	data, err := au.StudentRepository.LoginStudent(nim)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	claims := &student.JwtCustomClaimsStudent{
		ID:        data.ID,
		NIM:       data.NIM,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
		RoleId:    data.RoleId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 1)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	godotenv.Load()
	secretKeyStudent := os.Getenv("SECRET_KEY_STUDENT")
	newToken, err := token.SignedString([]byte(secretKeyStudent))
	if err != nil {
		return nil, 401, err
	}

	studentJWTResponse := student.StudentJWTResponse{
		ID:     data.ID,
		NIM:    data.NIM,
		RoleId: data.RoleId,
		Token:  newToken,
	}
	return &studentJWTResponse, http.StatusOK, nil
}
