package user

import (
	"context"
	"database/sql"
)

type mysqlStore struct {
	db *sql.DB
}

func NewStore(db *sql.DB) Repository {
	return &mysqlStore{db: db}
}

func (s *mysqlStore) List(ctx context.Context) ([]User, error) {
	rows, err := s.db.QueryContext(ctx,
		"SELECT user_id, name, email, language, timezone, privacy_mode, created_at FROM users",
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		u, err := scanUser(rows)
		if err != nil {
			return nil, err
		}
		users = append(users, *u)
	}
	return users, rows.Err()
}

func (s *mysqlStore) GetByID(ctx context.Context, id string) (*User, error) {
	rows, err := s.db.QueryContext(ctx,
		"SELECT user_id, name, email, language, timezone, privacy_mode, created_at FROM users WHERE user_id = ?",
		id,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if !rows.Next() {
		return nil, nil
	}
	return scanUser(rows)
}

func (s *mysqlStore) Create(ctx context.Context, u *User) error {
	_, err := s.db.ExecContext(ctx,
		"INSERT INTO users (user_id, name, email, language, timezone, privacy_mode) VALUES (?, ?, ?, ?, ?, ?)",
		u.UserID, u.Name, u.Email, u.Language, u.Timezone, u.PrivacyMode,
	)
	return err
}

func (s *mysqlStore) Update(ctx context.Context, u *User) error {
	_, err := s.db.ExecContext(ctx,
		"UPDATE users SET name = ?, email = ?, language = ?, timezone = ?, privacy_mode = ? WHERE user_id = ?",
		u.Name, u.Email, u.Language, u.Timezone, u.PrivacyMode, u.UserID,
	)
	return err
}

func (s *mysqlStore) Delete(ctx context.Context, id string) error {
	_, err := s.db.ExecContext(ctx, "DELETE FROM users WHERE user_id = ?", id)
	return err
}

type scanner interface {
	Scan(dest ...any) error
}

func scanUser(r scanner) (*User, error) {
	var u User
	var email, language, timezone, privacyMode, createdAt sql.NullString
	if err := r.Scan(&u.UserID, &u.Name, &email, &language, &timezone, &privacyMode, &createdAt); err != nil {
		return nil, err
	}
	if email.Valid {
		u.Email = &email.String
	}
	if language.Valid {
		u.Language = &language.String
	}
	if timezone.Valid {
		u.Timezone = &timezone.String
	}
	if privacyMode.Valid {
		u.PrivacyMode = &privacyMode.String
	}
	if createdAt.Valid {
		u.CreatedAt = &createdAt.String
	}
	return &u, nil
}
