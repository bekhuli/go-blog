package user

import (
	"github.com/bekhuli/go-blog/internal/common"
	"github.com/bekhuli/go-blog/pkg/auth"
	"github.com/gorilla/mux"
)

func RegisterUserRoutes(r *mux.Router, h *UserHandler) {
	// PUBLIC
	r.HandleFunc("/register", h.RegisterUser).Methods("POST")
	r.HandleFunc("/login", h.LoginUser).Methods("POST")

	// SECURED
	protected := r.PathPrefix("/").Subrouter()
	protected.Use(auth.JWTMiddleware(common.JWTEnv))

	protected.HandleFunc("/my-profile", h.GetUserByID).Methods("GET")
}
