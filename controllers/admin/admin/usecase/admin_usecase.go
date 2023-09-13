package usecase

import (
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/ardin2001/backend-pemilu/controllers/admin/admin"
	"github.com/ardin2001/backend-pemilu/controllers/admin/admin/repository"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

type AdminUsecaseInterface interface {
	GetById(id string) (*admin.AdminDetailResponse, int, error)
	LoginAdmin(username, password string) (*admin.AdminJWTResponse, int, error)
}

type AdminUsecaseStruct struct {
	AdminRepository repository.AdminRepositoryInterface
}

func NewAdminUsecase(AdminRepository repository.AdminRepositoryInterface) AdminUsecaseInterface {
	return &AdminUsecaseStruct{
		AdminRepository: AdminRepository,
	}
}

func (au *AdminUsecaseStruct) GetById(id string) (*admin.AdminDetailResponse, int, error) {
	admin, err := au.AdminRepository.GetById(id)
	if err != nil {
		return nil, http.StatusBadRequest, errors.New("failed to get detail data admin")
	}

	return admin, http.StatusOK, nil
}

func (au *AdminUsecaseStruct) LoginAdmin(username, password string) (*admin.AdminJWTResponse, int, error) {
	data, err := au.AdminRepository.LoginAdmin(username, password)
	if err != nil {
		return nil, http.StatusInternalServerError, errors.New("failed to login admin")
	}

	claims := &admin.JwtCustomClaimsAdmin{
		ID:        data.ID,
		Username:  data.Username,
		RoleId:    data.RoleId,
		CreatedAt: data.CreatedAt,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 1)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	godotenv.Load()
	secretKey := os.Getenv("SECRET_KEY_ADMIN")
	newToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return nil, 401, errors.New("failed to generate secret key jwt")
	}

	adminJWTResponse := admin.AdminJWTResponse{
		ID:       data.ID,
		Username: data.Username,
		RoleId:   data.RoleId,
		Token:    newToken,
	}
	return &adminJWTResponse, http.StatusOK, nil
}
