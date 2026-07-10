package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("dcp-secret-key")

func GenerateToken(userID uint64) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"user_id": userID,
			"exp":     time.Now().Add(time.Hour * 24).Unix(),
		})

	return token.SignedString(secretKey)
}
