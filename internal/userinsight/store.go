package userinsight

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

const selectCols = "SELECT insight_id, user_id, insight_type, title, description, johari_zone, status, recommended_action FROM user_insights"

func (s *mysqlStore) List(ctx context.Context) ([]UserInsight, error) {
	rows, err := s.db.QueryContext(ctx, selectCols)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanAll(rows)
}

func (s *mysqlStore) ListByUserID(ctx context.Context, userID string) ([]UserInsight, error) {
	rows, err := s.db.QueryContext(ctx, selectCols+" WHERE user_id = ?", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanAll(rows)
}

func (s *mysqlStore) GetByID(ctx context.Context, id string) (*UserInsight, error) {
	rows, err := s.db.QueryContext(ctx, selectCols+" WHERE insight_id = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if !rows.Next() {
		return nil, nil
	}
	return scanInsight(rows)
}

func (s *mysqlStore) Create(ctx context.Context, i *UserInsight) error {
	_, err := s.db.ExecContext(ctx,
		"INSERT INTO user_insights (insight_id, user_id, insight_type, title, description, johari_zone, status, recommended_action) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
		i.InsightID, i.UserID, i.InsightType, i.Title, i.Description, i.JohariZone, i.Status, i.RecommendedAction,
	)
	return err
}

func (s *mysqlStore) Update(ctx context.Context, i *UserInsight) error {
	_, err := s.db.ExecContext(ctx,
		"UPDATE user_insights SET user_id = ?, insight_type = ?, title = ?, description = ?, johari_zone = ?, status = ?, recommended_action = ? WHERE insight_id = ?",
		i.UserID, i.InsightType, i.Title, i.Description, i.JohariZone, i.Status, i.RecommendedAction, i.InsightID,
	)
	return err
}

func (s *mysqlStore) Delete(ctx context.Context, id string) error {
	_, err := s.db.ExecContext(ctx, "DELETE FROM user_insights WHERE insight_id = ?", id)
	return err
}

type scanner interface {
	Scan(dest ...any) error
}

func scanAll(rows *sql.Rows) ([]UserInsight, error) {
	var items []UserInsight
	for rows.Next() {
		i, err := scanInsight(rows)
		if err != nil {
			return nil, err
		}
		items = append(items, *i)
	}
	return items, rows.Err()
}

func scanInsight(r scanner) (*UserInsight, error) {
	var i UserInsight
	var insightType, title, description, johariZone, status, recommendedAction sql.NullString
	if err := r.Scan(&i.InsightID, &i.UserID, &insightType, &title, &description, &johariZone, &status, &recommendedAction); err != nil {
		return nil, err
	}
	if insightType.Valid {
		i.InsightType = &insightType.String
	}
	if title.Valid {
		i.Title = &title.String
	}
	if description.Valid {
		i.Description = &description.String
	}
	if johariZone.Valid {
		i.JohariZone = &johariZone.String
	}
	if status.Valid {
		i.Status = &status.String
	}
	if recommendedAction.Valid {
		i.RecommendedAction = &recommendedAction.String
	}
	return &i, nil
}
