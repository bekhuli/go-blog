package post

import (
	"github.com/bekhuli/go-blog/pkg/utils"
	"github.com/gorilla/mux"
	"net/http"
)

type PostHandler struct {
	service *Service
}

func NewPostHandler(service *Service) *PostHandler {
	return &PostHandler{service: service}
}

func (h *PostHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
	var dto CreatePostRequest
	if err := utils.BindJSON(r, &dto); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	createdPost, err := h.service.CreatePost(r.Context(), dto)
	if err != nil {
		utils.WriteError(w, http.StatusConflict, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, ToResponse(createdPost))
}

func (h *PostHandler) GetPostByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postID := vars["id"]

	post, err := h.service.GetPostByID(r.Context(), postID)
	if err != nil {
		utils.WriteError(w, http.StatusNotFound, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, ToResponse(post))
}

func (h *PostHandler) ListPosts(w http.ResponseWriter, r *http.Request) {
	authorID := r.URL.Query().Get("author_id")

	posts, err := h.service.ListPosts(r.Context(), authorID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, ToResponseList(posts))
}

func (h *PostHandler) UpdatePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postID := vars["id"]

	var dto UpdatePostRequest
	if err := utils.BindJSON(r, &dto); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	updatedPost, err := h.service.UpdatePost(r.Context(), &dto, postID)
	if err != nil {
		utils.WriteError(w, http.StatusForbidden, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, ToResponse(updatedPost))
}

func (h *PostHandler) DeletePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postID := vars["id"]

	resp, err := h.service.DeletePost(r.Context(), postID)
	if err != nil {
		utils.WriteError(w, http.StatusForbidden, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, resp)
}
