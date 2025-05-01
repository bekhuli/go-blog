package auth

import (
	"context"
	"errors"
	"fmt"
	"github.com/bekhuli/go-blog/internal/common"
	"github.com/bekhuli/go-blog/pkg/utils"
	"net/http"
	"strings"
)

type contextKey string

const UserKey contextKey = "userID"

func JWTMiddleware(cfg common.JWTConfig) func(handler http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				utils.WriteError(w, http.StatusUnauthorized, errors.New("authorization header required"))
				return
			}

			tokenString := strings.TrimPrefix(authHeader, "Bearer ")
			if tokenString == "" {
				utils.WriteError(w, http.StatusUnauthorized, errors.New("bearer token required"))
				return
			}

			claims, err := ParseJWT(tokenString, cfg)
			if err != nil {
				utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("invalid token: %v", err))
				return
			}

			ctx := context.WithValue(r.Context(), UserKey, claims.UserID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
