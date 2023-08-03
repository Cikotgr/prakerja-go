package admin

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type RequestLoginAdmin struct {
	Username string `json:"username" form:"username" validate:"required,min=5"`
	Password string `json:"password" form:"password" validate:"required,min=5"`
}

type JwtCustomClaimsAdmin struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	RoleId    int       `json:"role_id"`
	CreatedAt time.Time `json:"created_at"`
	jwt.RegisteredClaims
}
