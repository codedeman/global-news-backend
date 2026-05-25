package app

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

func (s *mysqlStore) List(ctx context.Context) ([]App, error) {
	rows, err := s.db.QueryContext(ctx, "SELECT id, name, COALESCE(description,'') FROM apps")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var apps []App
	for rows.Next() {
		var a App
		if err := rows.Scan(&a.ID, &a.Name, &a.Description); err != nil {
			return nil, err
		}
		apps = append(apps, a)
	}
	return apps, rows.Err()
}

func (s *mysqlStore) GetByID(ctx context.Context, id string) (*App, error) {
	var a App
	err := s.db.QueryRowContext(ctx,
		"SELECT id, name, COALESCE(description,'') FROM apps WHERE id = ?", id,
	).Scan(&a.ID, &a.Name, &a.Description)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &a, err
}

func (s *mysqlStore) Create(ctx context.Context, a *App) error {
	_, err := s.db.ExecContext(ctx,
		"INSERT INTO apps (id, name, description) VALUES (?, ?, ?)",
		a.ID, a.Name, a.Description,
	)
	return err
}

func (s *mysqlStore) Update(ctx context.Context, a *App) error {
	_, err := s.db.ExecContext(ctx,
		"UPDATE apps SET name = ?, description = ? WHERE id = ?",
		a.Name, a.Description, a.ID,
	)
	return err
}

func (s *mysqlStore) Delete(ctx context.Context, id string) error {
	_, err := s.db.ExecContext(ctx, "DELETE FROM apps WHERE id = ?", id)
	return err
}
