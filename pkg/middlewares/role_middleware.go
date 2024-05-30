package middlewares

import (
	"fmt"
	"net/http"

	"github.com/FrancoRutigliano/myMovies/pkg/helpers"
	authHelpers "github.com/FrancoRutigliano/myMovies/pkg/helpers/auth"
	"github.com/golang-jwt/jwt"
)

func RoleMiddleware(expRole string) func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			tokenStr := authHelpers.GetTokenFromRequest(r)
			if tokenStr == "" {
				helpers.SendCustom(w, http.StatusForbidden, "Unauthorized, token empty")
				return
			}

			t, err := authHelpers.ValidateToken(tokenStr)
			if err != nil {
				fmt.Println("Token validation error:", err)
				helpers.SendCustom(w, http.StatusForbidden, "Unauthorized, invalid token")
				return
			}

			claims, ok := t.Claims.(jwt.MapClaims)
			if !ok || !t.Valid {
				fmt.Println("Invalid claims or token is not valid")
				helpers.SendCustom(w, http.StatusForbidden, "Unauthorized")
				return
			}

			role, ok := claims["userRole"].(string)
			if !ok {
				fmt.Println("Role not found in token claims")
				helpers.SendCustom(w, http.StatusForbidden, "Unauthorized, role is not valid")
				return
			}

			if role != expRole {
				fmt.Println("Role mismatch: expected", expRole, "but got", role)
				helpers.SendCustom(w, http.StatusForbidden, "Unauthorized, role is not valid")
				return
			}

			next(w, r)
		}
	}
}
