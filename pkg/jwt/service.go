package jwt

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTService interface {
	GenerateToken(userID int) (string, error)
}

type jwtService struct {
	secretKey     string
	expireMinutes time.Duration
}

func NewJWTService() JWTService {
	expire := 15 // default
	if val := os.Getenv("JWT_EXPIRE_MINUTES"); val != "" {
		if parsed, err := time.ParseDuration(val + "m"); err == nil {
			expire = int(parsed.Minutes())
		}
	}
	return &jwtService{
		secretKey:     os.Getenv("JWT_SECRET"),
		expireMinutes: time.Duration(expire),
	}
}

func (j *jwtService) GenerateToken(userID int) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(j.expireMinutes * time.Minute).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.secretKey))
}
