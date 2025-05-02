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

func (s *Service) LoginUser(ctx context.Context, dto LoginRequest) (*User, error) {
	if err := s.validator.Validate(dto); err != nil {
		return nil, err
	}

	var user *User
	var err error

	if dto.Username != "" {
		user, err = s.repo.GetUserByUsername(ctx, dto.Username)
	} else {
		user, err = s.repo.GetUserByEmail(ctx, dto.Email)
	}

	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(dto.Password))
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	return user, nil
}
