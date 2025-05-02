package user

import (
	"context"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repo      Repository
	validator *Validator
}

func NewService(repo Repository, validator *Validator) *Service {
	return &Service{repo: repo, validator: validator}
}

func (s *Service) RegisterUser(ctx context.Context, dto RegisterRequest) (*User, error) {
	if err := s.validator.Validate(dto); err != nil {
		return nil, err
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)

	user := &User{
		ID:       uuid.New(),
		Username: dto.Username,
		Email:    dto.Email,
		Password: string(hashedPassword),
	}

	return s.repo.CreateUser(ctx, user)
}
