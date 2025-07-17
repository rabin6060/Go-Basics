package store

import (
	"context"
	"database/sql"
)

type Storage struct {
	Posts interface {
		createPost(context.Context,*Post) error
	}
	Users interface {
		createUser(context.Context,*User) error
	}
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Posts: &PostStore{db},
		Users: &UserStore{db},
	}
}
