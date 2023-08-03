package usecase

import (
	"net/http"
	"os"
	"time"

	"github.com/ardin2001/backend-pemilu/controllers/admin/admin"
	"github.com/ardin2001/backend-pemilu/controllers/admin/admin/repository"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

type AdminUsecaseInterface interface {
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

func (au *AdminUsecaseStruct) LoginAdmin(username, password string) (*admin.AdminJWTResponse, int, error) {
	data, err := au.AdminRepository.LoginAdmin(username, password)
	if err != nil {
		return nil, http.StatusInternalServerError, err
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
		return nil, 401, err
	}

	adminJWTResponse := admin.AdminJWTResponse{
		ID:       data.ID,
		Username: data.Username,
		RoleId:   data.RoleId,
		Token:    newToken,
	}
	return &adminJWTResponse, http.StatusOK, nil
}
