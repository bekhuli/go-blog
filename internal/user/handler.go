package user

import (
	"github.com/bekhuli/go-blog/pkg/utils"
	"net/http"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) RegisterUser(w http.ResponseWriter, r *http.Request) {
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
