package goal

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

func (s *mysqlStore) List(ctx context.Context) ([]Goal, error) {
	rows, err := s.db.QueryContext(ctx, "SELECT id, name FROM goals")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var goals []Goal
	for rows.Next() {
		var g Goal
		if err := rows.Scan(&g.ID, &g.Name); err != nil {
			return nil, err
		}
		goals = append(goals, g)
	}
	return goals, rows.Err()
}

func (s *mysqlStore) GetByID(ctx context.Context, id string) (*Goal, error) {
	var g Goal
	err := s.db.QueryRowContext(ctx, "SELECT id, name FROM goals WHERE id = ?", id).Scan(&g.ID, &g.Name)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &g, err
}

func (s *mysqlStore) Create(ctx context.Context, g *Goal) error {
	_, err := s.db.ExecContext(ctx, "INSERT INTO goals (id, name) VALUES (?, ?)", g.ID, g.Name)
	return err
}

func (s *mysqlStore) Update(ctx context.Context, g *Goal) error {
	_, err := s.db.ExecContext(ctx, "UPDATE goals SET name = ? WHERE id = ?", g.Name, g.ID)
	return err
}

func (s *mysqlStore) Delete(ctx context.Context, id string) error {
	_, err := s.db.ExecContext(ctx, "DELETE FROM goals WHERE id = ?", id)
	return err
}
