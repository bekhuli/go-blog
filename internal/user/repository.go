package user

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/lib/pq"
)

type Repository interface {
	CreateUser(ctx context.Context, user *User) (*User, error)
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	GetUserByUsername(ctx context.Context, username string) (*User, error)
	ExistsByEmail(ctx context.Context, email string) (bool, error)
	ExistsByUsername(ctx context.Context, username string) (bool, error)
}

type SQLRepository struct {
	db         *sql.DB
	userRoleID string
}

func NewUserRepository(db *sql.DB) (*SQLRepository, error) {
	const roleQuery = `SELECT id FROM roles WHERE role = 'user'`

	var roleID string
	err := db.QueryRow(roleQuery).Scan(&roleID)
	if err != nil {
		return nil, fmt.Errorf("failed to get role id: %w", err)
	}

	return &SQLRepository{db: db, userRoleID: roleID}, nil
}

func (r *SQLRepository) CreateUser(ctx context.Context, user *User) (*User, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %w", err)
	}

	const userQuery = `
		INSERT INTO users (id, username, email, password, created_at)
		VALUES ($1, $2, $3, $4, $5)
	`

	user.CreatedAt = time.Now().UTC()

	_, err = tx.ExecContext(
		ctx,
		userQuery,
		user.ID,
		user.Username,
		user.Email,
		user.Password,
		user.CreatedAt,
	)

	if err != nil {
		tx.Rollback()
		if isDuplicateError(err) {
			return nil, ErrUserExists
		}

		return nil, fmt.Errorf("user repository create: %w", err)
	}

	_, err = tx.ExecContext(ctx, `INSERT INTO user_role (user_id, role_id) VALUES ($1, $2)`, user.ID, r.userRoleID)
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("failed to add to user role: %w", err)
	}

	if err = tx.Commit(); err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return user, nil
}

func (r *SQLRepository) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	const query = `
		SELECT id, username, email, password, created_at
		FROM users
		WHERE email = $1
	`

	var user User
	err := r.db.QueryRowContext(ctx, query, email).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrUserNotFound
	}

	if err != nil {
		return nil, fmt.Errorf("get user by email: %w", err)
	}

	return &user, nil
}

func isDuplicateError(err error) bool {
	pgErr, ok := err.(*pq.Error)
	return ok && pgErr.Code == "23505"
}

func (r *SQLRepository) GetUserByUsername(ctx context.Context, username string) (*User, error) {
	const query = `
		SELECT id, username, email, password, created_at
		FROM users
		WHERE username = $1
	`

	var user User
	err := r.db.QueryRowContext(ctx, query, username).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrUserNotFound
	}

	if err != nil {
		return nil, fmt.Errorf("get user by username: %w", err)
	}

	return &user, nil
}

func (r *SQLRepository) ExistsByEmail(ctx context.Context, email string) (bool, error) {
	const query = `SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)`

	var exists bool
	err := r.db.QueryRowContext(ctx, query, email).Scan(&exists)

	if err != nil {
		return false, fmt.Errorf("check user exists by email: %w", err)
	}

	return exists, nil
}

func (r *SQLRepository) ExistsByUsername(ctx context.Context, username string) (bool, error) {
	const query = `SELECT EXISTS(SELECT 1 FROM users WHERE username = $1)`

	var exists bool
	err := r.db.QueryRowContext(ctx, query, username).Scan(&exists)

	if err != nil {
		return false, fmt.Errorf("check user exists by username: %w", err)
	}

	return exists, nil
}
