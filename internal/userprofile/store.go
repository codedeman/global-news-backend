package userprofile

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

const selectCols = "SELECT profile_id, user_id, age_range, gender_optional, country, persona, life_stage, updated_at FROM user_profiles"

func (s *mysqlStore) List(ctx context.Context) ([]UserProfile, error) {
	rows, err := s.db.QueryContext(ctx, selectCols)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanAll(rows)
}

func (s *mysqlStore) ListByUserID(ctx context.Context, userID string) ([]UserProfile, error) {
	rows, err := s.db.QueryContext(ctx, selectCols+" WHERE user_id = ?", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanAll(rows)
}

func (s *mysqlStore) GetByID(ctx context.Context, id string) (*UserProfile, error) {
	rows, err := s.db.QueryContext(ctx, selectCols+" WHERE profile_id = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if !rows.Next() {
		return nil, nil
	}
	return scanProfile(rows)
}

func (s *mysqlStore) Create(ctx context.Context, p *UserProfile) error {
	_, err := s.db.ExecContext(ctx,
		"INSERT INTO user_profiles (profile_id, user_id, age_range, gender_optional, country, persona, life_stage) VALUES (?, ?, ?, ?, ?, ?, ?)",
		p.ProfileID, p.UserID, p.AgeRange, p.GenderOptional, p.Country, p.Persona, p.LifeStage,
	)
	return err
}

func (s *mysqlStore) Update(ctx context.Context, p *UserProfile) error {
	_, err := s.db.ExecContext(ctx,
		"UPDATE user_profiles SET user_id = ?, age_range = ?, gender_optional = ?, country = ?, persona = ?, life_stage = ? WHERE profile_id = ?",
		p.UserID, p.AgeRange, p.GenderOptional, p.Country, p.Persona, p.LifeStage, p.ProfileID,
	)
	return err
}

func (s *mysqlStore) Delete(ctx context.Context, id string) error {
	_, err := s.db.ExecContext(ctx, "DELETE FROM user_profiles WHERE profile_id = ?", id)
	return err
}

type scanner interface {
	Scan(dest ...any) error
}

func scanAll(rows *sql.Rows) ([]UserProfile, error) {
	var items []UserProfile
	for rows.Next() {
		p, err := scanProfile(rows)
		if err != nil {
			return nil, err
		}
		items = append(items, *p)
	}
	return items, rows.Err()
}

func scanProfile(r scanner) (*UserProfile, error) {
	var p UserProfile
	var ageRange, gender, country, persona, lifeStage, updatedAt sql.NullString
	if err := r.Scan(&p.ProfileID, &p.UserID, &ageRange, &gender, &country, &persona, &lifeStage, &updatedAt); err != nil {
		return nil, err
	}
	if ageRange.Valid {
		p.AgeRange = &ageRange.String
	}
	if gender.Valid {
		p.GenderOptional = &gender.String
	}
	if country.Valid {
		p.Country = &country.String
	}
	if persona.Valid {
		p.Persona = &persona.String
	}
	if lifeStage.Valid {
		p.LifeStage = &lifeStage.String
	}
	if updatedAt.Valid {
		p.UpdatedAt = &updatedAt.String
	}
	return &p, nil
}
