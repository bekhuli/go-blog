package post

import "errors"

var (
	ErrPostNotFound = errors.New("post not found")
	ErrForbidden    = errors.New("forbidden")
	ErrUnauthorized = errors.New("not authorized")
)
