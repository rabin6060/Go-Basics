package store

import (
	"context"
	"database/sql"
)

type PostStore struct{
	db *sql.DB
}

func (s *PostStore) createPost(ctx context.Context) error {
	return nil
}