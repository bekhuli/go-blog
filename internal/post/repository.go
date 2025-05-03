package post

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Repository interface {
	CreatePost(ctx context.Context, post *Post) (*Post, error)
	GetPostByID(ctx context.Context, id string) (*Post, error)
	ListPosts(ctx context.Context, authorID string) ([]Post, error)
	UpdatePost(ctx context.Context, post *Post) (*Post, error)
	DeletePost(ctx context.Context, id, userID string) error
}

type SQLRepository struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) *SQLRepository {
	return &SQLRepository{db: db}
}

func (r *SQLRepository) CreatePost(ctx context.Context, post *Post) (*Post, error) {
	const query = `
		INSERT INTO posts (id, author_id, title, content, created_at)
		VALUES ($1, $2, $3, $4, $5)
	`

	post.ID = uuid.New()
	post.CreatedAt = time.Now().UTC()

	_, err := r.db.ExecContext(ctx, query,
		post.ID,
		post.AuthorID,
		post.Title,
		post.Content,
		post.CreatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("error inserting post: %v", err)
	}

	return post, nil
}

func (r *SQLRepository) GetPostByID(ctx context.Context, id string) (*Post, error) {
	const query = `
		SELECT id, author_id, title, content, created_at
		FROM posts
		WHERE id = $1
	`

	var post Post
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&post.ID,
		&post.AuthorID,
		&post.Title,
		&post.Content,
		&post.CreatedAt,
	)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrPostNotFound
	}

	if err != nil {
		return nil, fmt.Errorf("error getting post by ID: %w", err)
	}

	return &post, nil
}

func (r *SQLRepository) ListPosts(ctx context.Context, authorID string) ([]Post, error) {
	const query = `
		SELECT id, author_id, title, content, created_at
		FROM posts
		WHERE author_id = $1
		ORDER BY created_at DESC
	`

	rows, err := r.db.QueryContext(ctx, query, authorID)
	if err != nil {
		return nil, fmt.Errorf("error getting posts: %w", err)
	}

	defer rows.Close()

	var posts []Post
	for rows.Next() {
		var post Post
		if err := rows.Scan(&post.ID, &post.AuthorID, &post.Title, &post.Content, &post.CreatedAt); err != nil {
			return nil, fmt.Errorf("error scanning post: %w", err)
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func (r *SQLRepository) UpdatePost(ctx context.Context, post *Post) (*Post, error) {
	const query = `
		UPDATE posts
		SET title = $1, content = $2
		WHERE id = $3 AND author_id = $4
	`

	res, err := r.db.ExecContext(ctx, query,
		post.Title,
		post.Content,
		post.ID,
		post.AuthorID,
	)

	if err != nil {
		return nil, fmt.Errorf("error updating post: %w", err)
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return nil, ErrForbidden
	}

	return post, nil
}

func (r *SQLRepository) DeletePost(ctx context.Context, id, userID string) error {
	const query = `
		DELETE FROM posts WHERE id = $1 and author_id = $2
	`

	res, err := r.db.ExecContext(ctx, query, id, userID)
	if err != nil {
		return fmt.Errorf("error deleting post: %w", err)
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return ErrForbidden
	}

	return nil
}
