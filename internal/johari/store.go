package johari

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

const selectCols = "SELECT id, user_id, target_type, target_id, johari_zone, known_by_user, known_by_ai, confirmed_by_user, confidence_score FROM johari_understanding"

func (s *mysqlStore) List(ctx context.Context) ([]JohariUnderstanding, error) {
	rows, err := s.db.QueryContext(ctx, selectCols)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanAll(rows)
}

func (s *mysqlStore) ListByUserID(ctx context.Context, userID string) ([]JohariUnderstanding, error) {
	rows, err := s.db.QueryContext(ctx, selectCols+" WHERE user_id = ?", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanAll(rows)
}

func (s *mysqlStore) GetByID(ctx context.Context, id string) (*JohariUnderstanding, error) {
	rows, err := s.db.QueryContext(ctx, selectCols+" WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if !rows.Next() {
		return nil, nil
	}
	return scanJohari(rows)
}

func (s *mysqlStore) Create(ctx context.Context, j *JohariUnderstanding) error {
	_, err := s.db.ExecContext(ctx,
		"INSERT INTO johari_understanding (id, user_id, target_type, target_id, johari_zone, known_by_user, known_by_ai, confirmed_by_user, confidence_score) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)",
		j.ID, j.UserID, j.TargetType, j.TargetID, j.JohariZone, j.KnownByUser, j.KnownByAI, j.ConfirmedByUser, j.ConfidenceScore,
	)
	return err
}

func (s *mysqlStore) Update(ctx context.Context, j *JohariUnderstanding) error {
	_, err := s.db.ExecContext(ctx,
		"UPDATE johari_understanding SET user_id = ?, target_type = ?, target_id = ?, johari_zone = ?, known_by_user = ?, known_by_ai = ?, confirmed_by_user = ?, confidence_score = ? WHERE id = ?",
		j.UserID, j.TargetType, j.TargetID, j.JohariZone, j.KnownByUser, j.KnownByAI, j.ConfirmedByUser, j.ConfidenceScore, j.ID,
	)
	return err
}

func (s *mysqlStore) Delete(ctx context.Context, id string) error {
	_, err := s.db.ExecContext(ctx, "DELETE FROM johari_understanding WHERE id = ?", id)
	return err
}

type scanner interface {
	Scan(dest ...any) error
}

func scanAll(rows *sql.Rows) ([]JohariUnderstanding, error) {
	var items []JohariUnderstanding
	for rows.Next() {
		j, err := scanJohari(rows)
		if err != nil {
			return nil, err
		}
		items = append(items, *j)
	}
	return items, rows.Err()
}

func scanJohari(r scanner) (*JohariUnderstanding, error) {
	var j JohariUnderstanding
	var userID, targetType, targetID, johariZone sql.NullString
	var knownByUser, knownByAI, confirmedByUser sql.NullBool
	var confidenceScore sql.NullFloat64
	if err := r.Scan(&j.ID, &userID, &targetType, &targetID, &johariZone, &knownByUser, &knownByAI, &confirmedByUser, &confidenceScore); err != nil {
		return nil, err
	}
	if userID.Valid {
		j.UserID = &userID.String
	}
	if targetType.Valid {
		j.TargetType = &targetType.String
	}
	if targetID.Valid {
		j.TargetID = &targetID.String
	}
	if johariZone.Valid {
		j.JohariZone = &johariZone.String
	}
	if knownByUser.Valid {
		j.KnownByUser = &knownByUser.Bool
	}
	if knownByAI.Valid {
		j.KnownByAI = &knownByAI.Bool
	}
	if confirmedByUser.Valid {
		j.ConfirmedByUser = &confirmedByUser.Bool
	}
	if confidenceScore.Valid {
		j.ConfidenceScore = &confidenceScore.Float64
	}
	return &j, nil
}
