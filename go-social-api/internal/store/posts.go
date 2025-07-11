package store

import (
	"context"
	"database/sql"
	"github.com/lib/pq"
)

type Post struct {
	ID        int64    `json:"id"`
	Title     string   `json:"title"`
	Content   string   `json:"content"`
	Tags      []string `json:"tags"`
	UserID    int64    `json:"user_id"`
	CreatedAt string   `json:"created_at"`
	UpdatedAt string   `json:"updated_at"`
}
type PostStore struct {
	db *sql.DB
}

func (s *PostStore) Create(ctx context.Context, post *Post) error {
	query := `
		INSERT INTO post (title,content, tags,user_id)
		VALUES ($1,$2,$3,$4) RETURNING id,created_at,updated_at
	`
	err := s.db.QueryRowContext(
		ctx,
		query,
		post.Title, post.Content, pq.Array(post.Tags), post.UserID).
		Scan(&post.ID, &post.CreatedAt, &post.UpdatedAt)

	if err != nil {
		return err
	}
	return nil
}
