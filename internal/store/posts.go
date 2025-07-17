package store

import (
	"context"
	"database/sql"
)

//define schema
type Post struct {
	ID 		int64		`json:"id"`
	Title 	string		`json:"title"`
	Content string		`json:"content"`
	UserId 	int64		`json:"user_ID"`
	Tags   	[]string	`json:"tags"`
	CreatedAt string	`json:"created_At"`
	UpdatedAt string	`json:"updated_At"`
}

type PostStore struct{
	db *sql.DB
}

func (s *PostStore) createPost(ctx context.Context,post *Post) error {
	query := `INSERT INTO post (content,title,user_ID,tags) VALUES($1,$2,$3,$4) RETURNING 
				id,created_At `
	err:=s.db.QueryRowContext(ctx,query,
	post.Content,
	post.Title,
	post.UserId,
	post.Tags,
		).Scan(post.ID,post.CreatedAt)
	
	if err!=nil{
		return  err
	}
	return  nil
}