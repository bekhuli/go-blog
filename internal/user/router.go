package user

import "github.com/gorilla/mux"

func RegisterRoutes(r *mux.Router, h *Handler) {
	r.HandleFunc("/register", h.RegisterUser).Methods("POST")
	r.HandleFunc("/login", h.LoginUser).Methods("POST")
}
