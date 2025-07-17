package store

import (
	"context"
	"database/sql"
)

type User struct{
	Username string `json:"username"`
	Email 	 string `json:"email"`
	Password string `json:"-"`
	CreatedAt string `json:"created_At"`
}

type UserStore struct{
	db *sql.DB
}

func (s *UserStore) createUser(ctx context.Context,user *User) error {
	query := `INSERT INTO user (username,email,password) VALUES($1,$2,$3) RETURNING 
				username,created_At `
	err:=s.db.QueryRowContext(ctx,query,
	user.Username,
	user.Email,
	user.Password,
		).Scan(user.Username,user.CreatedAt)
	
	if err!=nil{
		return  err
	}
	return  nil
}