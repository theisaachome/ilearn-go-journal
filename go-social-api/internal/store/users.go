package store

import (
	"context"
	"database/sql"
)

type User struct {
	ID        int64  `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	PassWord  string `json:"-"`
	CreatedAt string `json:"created_at"`
}
type UserStore struct {
	db *sql.DB
}

func (s *UserStore) Create(ctx context.Context, user *User) error {
	query := `
	  INSERT INTO users (Username, email, password)
	  VALUES ($1, $2, $3) RETURNING ID,NAME,CREATED_AT
    `
	err := s.db.QueryRowContext(ctx, query,
		user.Username,
		user.Email,
		user.PassWord,
	).Scan(&user.ID, &user.Username, &user.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}
