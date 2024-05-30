package middlewares

import (
	"net/http"

	"github.com/FrancoRutigliano/myMovies/pkg/helpers"
	authHelpers "github.com/FrancoRutigliano/myMovies/pkg/helpers/auth"
	"github.com/golang-jwt/jwt"
)

func RoleMiddleware(expRole string) func(http.Handler) http.HandlerFunc {
	return func(next http.Handler) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			tokenStr := authHelpers.GetTokenFromRequest(r)
			if tokenStr == "" {
				helpers.SendCustom(w, http.StatusForbidden, "Unauthorized")
				return
			}

			// validar la signature del token para saber que el token no ha sido alterado
			t, err := authHelpers.ValidateToken(tokenStr)
			if err != nil {
				helpers.SendCustom(w, http.StatusForbidden, "Unauthorized")
				return
			}

			claims, ok := t.Claims.(jwt.MapClaims)
			if !ok || !t.Valid {
				helpers.SendCustom(w, http.StatusForbidden, "Unauthorized")
				return
			}

			role := claims["userRole"].(string)
			if role == "admin" {
				next.ServeHTTP(w, r)
			}
			// si el rol no coincide
			if role != expRole {
				helpers.SendCustom(w, http.StatusForbidden, "Unauthorized")
				return
			}

			// pasar el request al manejador siguiente
			next.ServeHTTP(w, r)
		}
	}
}
