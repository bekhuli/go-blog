package post

import "time"

type CreatePostRequest struct {
	Title   string `json:"title" validate:"required,min=1"`
	Content string `json:"content" validate:"required,min=1"`
}

type UpdatePostRequest struct {
	Title   string `json:"title" validate:"required,min=1"`
	Content string `json:"content" validate:"required,min=1"`
}

type PostResponse struct {
	ID        string    `json:"id"`
	AuthorID  string    `json:"author_id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

type DeletePostResponse struct {
	Message string `json:"message"`
}

func ToResponse(p *Post) *PostResponse {
	return &PostResponse{
		ID:        p.ID.String(),
		AuthorID:  p.AuthorID.String(),
		Title:     p.Title,
		Content:   p.Content,
		CreatedAt: p.CreatedAt,
	}
}
