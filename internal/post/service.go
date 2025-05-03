package post

import (
	"context"
	"github.com/bekhuli/go-blog/pkg/auth"

	"github.com/google/uuid"
)

type Service struct {
	repo      Repository
	validator *Validator
}

func NewPostService(repo Repository, validator *Validator) *Service {
	return &Service{repo: repo, validator: validator}
}

func (s *Service) CreatePost(ctx context.Context, dto CreatePostRequest) (*Post, error) {
	if err := s.validator.Validate(dto); err != nil {
		return nil, err
	}

	userID, ok := ctx.Value(auth.UserKey).(uuid.UUID)
	if !ok {
		return nil, ErrUnauthorized
	}

	post := &Post{
		ID:       uuid.New(),
		AuthorID: userID,
		Title:    dto.Title,
		Content:  dto.Content,
	}

	return s.repo.CreatePost(ctx, post)
}

func (s *Service) GetPostByID(ctx context.Context, id string) (*Post, error) {
	return s.repo.GetPostByID(ctx, id)
}

func (s *Service) ListPosts(ctx context.Context, authorID string) ([]Post, error) {
	return s.repo.ListPosts(ctx, authorID)
}

func (s *Service) UpdatePost(ctx context.Context, dto *UpdatePostRequest, postID string) (*Post, error) {
	if err := s.validator.Validate(dto); err != nil {
		return nil, err
	}

	userID, ok := ctx.Value(auth.UserKey).(string)
	if !ok {
		return nil, ErrUnauthorized
	}

	existingPost, err := s.repo.GetPostByID(ctx, postID)
	if err != nil {
		return nil, err
	}

	if existingPost.AuthorID.String() != userID {
		return nil, ErrForbidden
	}

	existingPost.Title = dto.Title
	existingPost.Content = dto.Content

	updatedPost, err := s.repo.UpdatePost(ctx, existingPost)
	if err != nil {
		return nil, err
	}

	return updatedPost, nil
}

func (s *Service) DeletePost(ctx context.Context, id string) (*DeletePostResponse, error) {
	userID, ok := ctx.Value(auth.UserKey).(string)
	if !ok {
		return nil, ErrUnauthorized
	}

	err := s.repo.DeletePost(ctx, id, userID)
	if err != nil {
		return nil, err
	}

	return &DeletePostResponse{
		Message: "Post successfully deleted",
	}, nil
}
