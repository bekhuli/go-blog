package routes

import (
	"database/sql"
	"github.com/bekhuli/go-blog/internal/user"
	"github.com/gorilla/mux"
)

func InitRouter(db *sql.DB) *mux.Router {
	r := mux.NewRouter()

	api := r.PathPrefix("/api/v1").Subrouter()

	userRepo := user.NewSQLRepository(db)
	userValidator := user.NewValidator()
	userService := user.NewService(userRepo, userValidator)
	userHandler := user.NewHandler(userService)

	user.RegisterRoutes(api, userHandler)

	return r
}
