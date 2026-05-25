package userbehavior

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

const selectCols = "SELECT signal_id, user_id, signal_type, raw_event_ref, related_category_id, intensity, frequency, confidence_score FROM user_behavior_signals"

func (s *mysqlStore) List(ctx context.Context) ([]UserBehaviorSignal, error) {
	rows, err := s.db.QueryContext(ctx, selectCols)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanAll(rows)
}

func (s *mysqlStore) ListByUserID(ctx context.Context, userID string) ([]UserBehaviorSignal, error) {
	rows, err := s.db.QueryContext(ctx, selectCols+" WHERE user_id = ?", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanAll(rows)
}

func (s *mysqlStore) GetByID(ctx context.Context, id string) (*UserBehaviorSignal, error) {
	rows, err := s.db.QueryContext(ctx, selectCols+" WHERE signal_id = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if !rows.Next() {
		return nil, nil
	}
	return scanSignal(rows)
}

func (s *mysqlStore) Create(ctx context.Context, sig *UserBehaviorSignal) error {
	_, err := s.db.ExecContext(ctx,
		"INSERT INTO user_behavior_signals (signal_id, user_id, signal_type, raw_event_ref, related_category_id, intensity, frequency, confidence_score) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
		sig.SignalID, sig.UserID, sig.SignalType, sig.RawEventRef, sig.RelatedCategoryID, sig.Intensity, sig.Frequency, sig.ConfidenceScore,
	)
	return err
}

func (s *mysqlStore) Update(ctx context.Context, sig *UserBehaviorSignal) error {
	_, err := s.db.ExecContext(ctx,
		"UPDATE user_behavior_signals SET user_id = ?, signal_type = ?, raw_event_ref = ?, related_category_id = ?, intensity = ?, frequency = ?, confidence_score = ? WHERE signal_id = ?",
		sig.UserID, sig.SignalType, sig.RawEventRef, sig.RelatedCategoryID, sig.Intensity, sig.Frequency, sig.ConfidenceScore, sig.SignalID,
	)
	return err
}

func (s *mysqlStore) Delete(ctx context.Context, id string) error {
	_, err := s.db.ExecContext(ctx, "DELETE FROM user_behavior_signals WHERE signal_id = ?", id)
	return err
}

type scanner interface {
	Scan(dest ...any) error
}

func scanAll(rows *sql.Rows) ([]UserBehaviorSignal, error) {
	var items []UserBehaviorSignal
	for rows.Next() {
		sig, err := scanSignal(rows)
		if err != nil {
			return nil, err
		}
		items = append(items, *sig)
	}
	return items, rows.Err()
}

func scanSignal(r scanner) (*UserBehaviorSignal, error) {
	var sig UserBehaviorSignal
	var signalType, rawEventRef, relatedCategoryID sql.NullString
	var intensity, confidenceScore sql.NullFloat64
	var frequency sql.NullInt64
	if err := r.Scan(&sig.SignalID, &sig.UserID, &signalType, &rawEventRef, &relatedCategoryID, &intensity, &frequency, &confidenceScore); err != nil {
		return nil, err
	}
	if signalType.Valid {
		sig.SignalType = &signalType.String
	}
	if rawEventRef.Valid {
		sig.RawEventRef = &rawEventRef.String
	}
	if relatedCategoryID.Valid {
		sig.RelatedCategoryID = &relatedCategoryID.String
	}
	if intensity.Valid {
		sig.Intensity = &intensity.Float64
	}
	if frequency.Valid {
		v := int(frequency.Int64)
		sig.Frequency = &v
	}
	if confidenceScore.Valid {
		sig.ConfidenceScore = &confidenceScore.Float64
	}
	return &sig, nil
}
