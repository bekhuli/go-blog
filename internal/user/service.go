package user

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repo      Repository
	validator *Validator
}

func NewUserService(repo Repository, validator *Validator) *Service {
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

	if exists, err := s.repo.ExistsByUsername(ctx, user.Username); err != nil {
		return nil, fmt.Errorf("check if username exists: %w", err)
	} else if exists {
		return nil, errors.New("username already exists")
	}

	if exists, err := s.repo.ExistsByUsername(ctx, user.Username); err != nil {
		return nil, fmt.Errorf("check if email exists: %w", err)
	} else if exists {
		return nil, errors.New("email already exists")
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
		return nil, ErrInvalidPassword
	}

	return user, nil
}

func (s *Service) GetUserByID(ctx context.Context, id string) (*User, error) {
	return s.repo.GetUserByID(ctx, id)
}
