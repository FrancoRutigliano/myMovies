package authHelpers

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
)

func CreateJwt(secret []byte, userRole, userEmail string) (string, error) {
	// expira en 24 horas == 1 día
	expiration := time.Now().Add(time.Hour * 24).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
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

func GetTokenFromRequest(r *http.Request) string {
	tokenAuth := r.Header.Get("Authorization")

	// verificamos si el encabezado no esta vacio
	// y el prefijo Bearer
	if len(tokenAuth) == 0 && strings.HasPrefix(tokenAuth, "Bearer") {
		return strings.TrimPrefix(tokenAuth, "Bearer ")
	}
	// si el encabezado authorization no tiene nada o no comienza con bearer devolvemos vacio
	return ""
}

// Validar un token JWT implica verificar su integridad y autenticidad para asegurarse
// de que no ha sido alterado y de que fue emitido por una fuente confiable
func ValidateToken(t string) (*jwt.Token, error) {
	// utilizamos esta funcion para analizar al token
	return jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		// verificamos que la firma del token sea valida
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token")
		}

		return []byte(viper.GetString("JWT_SECRET")), nil
	})
}
