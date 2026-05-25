package screen

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

func (s *mysqlStore) List(ctx context.Context) ([]Screen, error) {
	rows, err := s.db.QueryContext(ctx, "SELECT id, app_id, name FROM screens")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var screens []Screen
	for rows.Next() {
		var sc Screen
		if err := rows.Scan(&sc.ID, &sc.AppID, &sc.Name); err != nil {
			return nil, err
		}
		screens = append(screens, sc)
	}
	return screens, rows.Err()
}

func (s *mysqlStore) ListByApp(ctx context.Context, appID string) ([]Screen, error) {
	rows, err := s.db.QueryContext(ctx,
		"SELECT id, app_id, name FROM screens WHERE app_id = ?", appID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var screens []Screen
	for rows.Next() {
		var sc Screen
		if err := rows.Scan(&sc.ID, &sc.AppID, &sc.Name); err != nil {
			return nil, err
		}
		screens = append(screens, sc)
	}
	return screens, rows.Err()
}

func (s *mysqlStore) GetByID(ctx context.Context, id string) (*Screen, error) {
	var sc Screen
	err := s.db.QueryRowContext(ctx,
		"SELECT id, app_id, name FROM screens WHERE id = ?", id,
	).Scan(&sc.ID, &sc.AppID, &sc.Name)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &sc, err
}

func (s *mysqlStore) Create(ctx context.Context, sc *Screen) error {
	_, err := s.db.ExecContext(ctx,
		"INSERT INTO screens (id, app_id, name) VALUES (?, ?, ?)",
		sc.ID, sc.AppID, sc.Name,
	)
	return err
}

func (s *mysqlStore) Update(ctx context.Context, sc *Screen) error {
	_, err := s.db.ExecContext(ctx,
		"UPDATE screens SET app_id = ?, name = ? WHERE id = ?",
		sc.AppID, sc.Name, sc.ID,
	)
	return err
}

func (s *mysqlStore) Delete(ctx context.Context, id string) error {
	_, err := s.db.ExecContext(ctx, "DELETE FROM screens WHERE id = ?", id)
	return err
}
