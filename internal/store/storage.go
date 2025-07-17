package store

import (
	"context"
	"database/sql"
)

type Storage struct {
	Posts interface {
		createPost(context.Context) error
	}
	Users interface {
		createUser(context.Context) error
	}
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Posts: &PostStore{db},
		Users: &UserStore{db},
	}
}
