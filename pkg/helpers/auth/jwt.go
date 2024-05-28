package authHelpers

import (
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateJwt(secret []byte, userId, userRole, userEmail string) (string, error) {
	// expira en 24 horas == 1 día
	expiration := time.Now().Add(time.Hour * 24).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":    userId,
		"email":     userEmail,
		"userRole":  userRole,
		"expiredAt": expiration,
	})

	// firmar el token
	tokenStr, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}
