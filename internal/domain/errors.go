package domain

import "errors"

var (
	ErrConnectPostgreSQL = errors.New("all attempts are exceeded, unable to connect to PostgreSQL")
	ErrInvalidInput      = errors.New("invalid input body")
)
