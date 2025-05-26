package user

import (
	"errors"
	"net/http"

	"github.com/bekhuli/go-blog/internal/common"
	"github.com/bekhuli/go-blog/pkg/auth"
	"github.com/bekhuli/go-blog/pkg/utils"
)

type UserHandler struct {
	service *Service
}

func NewUserHandler(service *Service) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var dto RegisterRequest
	if err := utils.BindJSON(r, &dto); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	user, err := h.service.RegisterUser(r.Context(), dto)
	if err != nil {
		utils.WriteError(w, http.StatusConflict, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, ToResponse(user))
}

func (h *UserHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	var dto LoginRequest
	if err := utils.BindJSON(r, &dto); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	user, err := h.service.LoginUser(r.Context(), dto)
	if err != nil {
		utils.WriteError(w, http.StatusUnauthorized, err)
		return
	}

	token, err := auth.GenerateJWT(common.JWTEnv, user.ID, user.Username)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]string{"token": token})
}

func (h *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	userID, ok := auth.GetUserIDFromContext(r.Context())
	if !ok {
		utils.WriteError(w, http.StatusUnauthorized, errors.New("unauthorized"))
		return
	}

	user, err := h.service.GetUserByID(r.Context(), userID.String())
	if err != nil {
		utils.WriteError(w, http.StatusNotFound, err)
	}

	utils.WriteJSON(w, http.StatusOK, ToResponse(user))
}
