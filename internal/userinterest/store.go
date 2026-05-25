package userinterest

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

const selectCols = "SELECT user_interest_id, user_id, category_id, interest_strength, source, confidence_score, user_confirmed FROM user_interests"

func (s *mysqlStore) List(ctx context.Context) ([]UserInterest, error) {
	rows, err := s.db.QueryContext(ctx, selectCols)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanAll(rows)
}

func (s *mysqlStore) ListByUserID(ctx context.Context, userID string) ([]UserInterest, error) {
	rows, err := s.db.QueryContext(ctx, selectCols+" WHERE user_id = ?", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanAll(rows)
}

func (s *mysqlStore) GetByID(ctx context.Context, id string) (*UserInterest, error) {
	rows, err := s.db.QueryContext(ctx, selectCols+" WHERE user_interest_id = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if !rows.Next() {
		return nil, nil
	}
	return scanInterest(rows)
}

func (s *mysqlStore) Create(ctx context.Context, i *UserInterest) error {
	_, err := s.db.ExecContext(ctx,
		"INSERT INTO user_interests (user_interest_id, user_id, category_id, interest_strength, source, confidence_score, user_confirmed) VALUES (?, ?, ?, ?, ?, ?, ?)",
		i.UserInterestID, i.UserID, i.CategoryID, i.InterestStrength, i.Source, i.ConfidenceScore, i.UserConfirmed,
	)
	return err
}

func (s *mysqlStore) Update(ctx context.Context, i *UserInterest) error {
	_, err := s.db.ExecContext(ctx,
		"UPDATE user_interests SET user_id = ?, category_id = ?, interest_strength = ?, source = ?, confidence_score = ?, user_confirmed = ? WHERE user_interest_id = ?",
		i.UserID, i.CategoryID, i.InterestStrength, i.Source, i.ConfidenceScore, i.UserConfirmed, i.UserInterestID,
	)
	return err
}

func (s *mysqlStore) Delete(ctx context.Context, id string) error {
	_, err := s.db.ExecContext(ctx, "DELETE FROM user_interests WHERE user_interest_id = ?", id)
	return err
}

type scanner interface {
	Scan(dest ...any) error
}

func scanAll(rows *sql.Rows) ([]UserInterest, error) {
	var items []UserInterest
	for rows.Next() {
		i, err := scanInterest(rows)
		if err != nil {
			return nil, err
		}
		items = append(items, *i)
	}
	return items, rows.Err()
}

func scanInterest(r scanner) (*UserInterest, error) {
	var i UserInterest
	var categoryID, source sql.NullString
	var interestStrength, confidenceScore sql.NullFloat64
	var userConfirmed sql.NullString
	if err := r.Scan(&i.UserInterestID, &i.UserID, &categoryID, &interestStrength, &source, &confidenceScore, &userConfirmed); err != nil {
		return nil, err
	}
	if categoryID.Valid {
		i.CategoryID = &categoryID.String
	}
	if interestStrength.Valid {
		i.InterestStrength = &interestStrength.Float64
	}
	if source.Valid {
		i.Source = &source.String
	}
	if confidenceScore.Valid {
		i.ConfidenceScore = &confidenceScore.Float64
	}
	if userConfirmed.Valid {
		i.UserConfirmed = &userConfirmed.String
	}
	return &i, nil
}
