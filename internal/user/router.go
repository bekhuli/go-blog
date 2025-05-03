package user

import "github.com/gorilla/mux"

func RegisterUserRoutes(r *mux.Router, h *UserHandler) {
	r.HandleFunc("/register", h.RegisterUser).Methods("POST")
	r.HandleFunc("/login", h.LoginUser).Methods("POST")
}
