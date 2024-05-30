package authHelpers

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateJwt(secret []byte, userRole, userEmail string) (string, error) {
	// Expira en 24 horas == 1 día
	expiration := time.Now().Add(time.Hour * 24).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":     userEmail,
		"userRole":  userRole,
		"expiredAt": expiration,
	})

	// Firmar el token
	tokenStr, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

func GetTokenFromRequest(r *http.Request) string {
	tokenAuth := r.Header.Get("Authorization")

	// verificamos si el encabezado no esta vacio
	// y el prefijo Bearer
	if tokenAuth != "" {
		return tokenAuth
	}
	return ""
}

// Validar un token JWT implica verificar su integridad y autenticidad para asegurarse
// de que no ha sido alterado y de que fue emitido por una fuente confiable
// Función para validar el token
func ValidateToken(t string) (*jwt.Token, error) {
	// Parseamos el token
	t = strings.TrimPrefix(t, "Bearer ")

	token, err := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}
