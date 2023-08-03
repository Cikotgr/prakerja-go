package student

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type RequestLoginStudent struct {
	NIM string `json:"nim" form:"nim" validate:"required,number,min=12"`
}

type JwtCustomClaimsStudent struct {
	ID        string    `json:"id"`
	NIM       string    `json:"nim"`
	RoleId    int       `json:"role_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	jwt.RegisteredClaims
}
