package post

import (
	"github.com/bekhuli/go-blog/internal/common"
	"github.com/bekhuli/go-blog/pkg/auth"

	"github.com/gorilla/mux"
)

func RegisterPostRoutes(r *mux.Router, h *PostHandler) {
	// PUBLIC
	r.HandleFunc("/posts", h.ListPosts).Methods("GET")
	r.HandleFunc("/posts/{id}", h.GetPostByID).Methods("GET")

	// SECURED
	protected := r.NewRoute().Subrouter()
	protected.Use(auth.JWTMiddleware(common.JWTEnv))

	protected.HandleFunc("/posts", h.CreatePost).Methods("POST")
	protected.HandleFunc("/posts/{id}", h.UpdatePost).Methods("PUT")
	protected.HandleFunc("/posts/{id}", h.DeletePost).Methods("DELETE")
}
