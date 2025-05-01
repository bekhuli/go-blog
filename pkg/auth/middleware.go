package auth

import (
	"context"
	"errors"
	"fmt"
	"github.com/bekhuli/go-blog/internal/common/config"
	"github.com/bekhuli/go-blog/pkg/utils"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const UserKey contextKey = "userID"

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg := config.JWTEnv

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			utils.WriteError(w, http.StatusBadRequest, errors.New("no token provided"))
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(cfg.JWTSecret), nil
		})

		if err != nil || !token.Valid {
			utils.WriteError(w, http.StatusUnauthorized, errors.New("invalid token"))
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			utils.WriteError(w, http.StatusUnauthorized, errors.New("invalid token"))
			return
		}

		UserIDFloat, ok := claims["userID"].(float64)
		if !ok {
			utils.WriteError(w, http.StatusUnauthorized, errors.New("invalid userID in token"))
			return
		}

		userID := int64(UserIDFloat)
		ctx := context.WithValue(r.Context(), UserKey, userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
