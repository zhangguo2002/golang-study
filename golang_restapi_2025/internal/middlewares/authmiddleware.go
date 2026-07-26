package middlewares

import (
	"context"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/zhangguo2002/golangrestapi/internal/auth"
	"github.com/zhangguo2002/golangrestapi/internal/utils"
)

// creates a custom type for context key to avoid collision
type contextKey string

// constant used in storing our user claims
const UserClaimsKey contextKey = "claims"

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//retrieves the Authorization header from the request (postman/web/mobile)
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			utils.RespondWithError(w, http.StatusUnauthorized, "No token provided")
			return
		}
		//strips the Bearer from the Bearer Token
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		claims := &auth.Claims{}
		//Parse the token and validate it
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			//we provide our key from the environment variable and validate it against the token from the request
			return []byte(os.Getenv("JWT_SECRET_KEY")), nil
		})
		//handle the validation error
		if err != nil {
			//handling likely tampered token
			if err == jwt.ErrSignatureInvalid {
				utils.RespondWithError(w, http.StatusBadRequest, "invalid token signature")
				return
			}
			//handles any other parsing error e.g expired malformed etc
			utils.RespondWithError(w, http.StatusBadRequest, "invalid token")
			return
		}
		//if token is valid ,store the claims in the request context
		if token.Valid {
			ctx := context.WithValue(r.Context(), UserClaimsKey, claims)
			r = r.WithContext(ctx) //replace request context with the new one
			next.ServeHTTP(w, r)   //calls the next handler ,with the updated request
		} else {
			utils.RespondWithError(w, http.StatusBadRequest, "invalid token")
		}
	})
}
