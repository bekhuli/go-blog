package routes

import (
	"database/sql"

	"github.com/bekhuli/go-blog/internal/post"
	"github.com/bekhuli/go-blog/internal/user"

	"github.com/gorilla/mux"
)

func InitRouter(db *sql.DB) *mux.Router {
	r := mux.NewRouter()

	api := r.PathPrefix("/api/v1").Subrouter()

	// --- USER ---
	userRepo := user.NewUserRepository(db)
	userValidator := user.NewUserValidator()
	userService := user.NewUserService(userRepo, userValidator)
	userHandler := user.NewUserHandler(userService)

	user.RegisterUserRoutes(api, userHandler)

	// --- POST ---
	postRepo := post.NewPostRepository(db)
	postValidator := post.NewPostValidator()
	postService := post.NewPostService(postRepo, postValidator)
	postHandler := post.NewPostHandler(postService)

	post.RegisterPostRoutes(api, postHandler)

	return r
}
