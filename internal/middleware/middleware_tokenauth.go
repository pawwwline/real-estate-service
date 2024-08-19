package middleware

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"real-estate-service/api/generated"
	"strings"
)

const (
	BearerAuthScopes = generated.BearerAuthScopes
	TokenSign        = "b99f6a9b74321b8b4f4c73e3de004ad7a3bd78f3482e93c8f4a596a6b09f2208c"
)

func TokenAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			http.Error(w, "Authorization header is missing", http.StatusUnauthorized)
			return
		}

		tokenParts := strings.Split(token, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			http.Error(w, "Invalid token format", http.StatusUnauthorized)
			return
		}

		jwtToken := tokenParts[1]
		claims := jwt.MapClaims{}
		t, err := jwt.ParseWithClaims(jwtToken, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(TokenSign), nil
		})

		if err != nil || !t.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		user_type, ok := claims["user_type"].(string)
		if !ok {
			http.Error(w, "Invalid token payload", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), BearerAuthScopes, user_type)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
