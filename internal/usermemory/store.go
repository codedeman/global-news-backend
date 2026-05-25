package usermemory

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

const selectCols = "SELECT memory_id, user_id, memory_type, title, content, source, confidence_score, sensitivity_level, created_at FROM user_memories"

func (s *mysqlStore) List(ctx context.Context) ([]UserMemory, error) {
	rows, err := s.db.QueryContext(ctx, selectCols)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanAll(rows)
}

func (s *mysqlStore) ListByUserID(ctx context.Context, userID string) ([]UserMemory, error) {
	rows, err := s.db.QueryContext(ctx, selectCols+" WHERE user_id = ?", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanAll(rows)
}

func (s *mysqlStore) GetByID(ctx context.Context, id string) (*UserMemory, error) {
	rows, err := s.db.QueryContext(ctx, selectCols+" WHERE memory_id = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if !rows.Next() {
		return nil, nil
	}
	return scanMemory(rows)
}

func (s *mysqlStore) Create(ctx context.Context, m *UserMemory) error {
	_, err := s.db.ExecContext(ctx,
		"INSERT INTO user_memories (memory_id, user_id, memory_type, title, content, source, confidence_score, sensitivity_level) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
		m.MemoryID, m.UserID, m.MemoryType, m.Title, m.Content, m.Source, m.ConfidenceScore, m.SensitivityLevel,
	)
	return err
}

func (s *mysqlStore) Update(ctx context.Context, m *UserMemory) error {
	_, err := s.db.ExecContext(ctx,
		"UPDATE user_memories SET user_id = ?, memory_type = ?, title = ?, content = ?, source = ?, confidence_score = ?, sensitivity_level = ? WHERE memory_id = ?",
		m.UserID, m.MemoryType, m.Title, m.Content, m.Source, m.ConfidenceScore, m.SensitivityLevel, m.MemoryID,
	)
	return err
}

func (s *mysqlStore) Delete(ctx context.Context, id string) error {
	_, err := s.db.ExecContext(ctx, "DELETE FROM user_memories WHERE memory_id = ?", id)
	return err
}

type scanner interface {
	Scan(dest ...any) error
}

func scanAll(rows *sql.Rows) ([]UserMemory, error) {
	var items []UserMemory
	for rows.Next() {
		m, err := scanMemory(rows)
		if err != nil {
			return nil, err
		}
		items = append(items, *m)
	}
	return items, rows.Err()
}

func scanMemory(r scanner) (*UserMemory, error) {
	var m UserMemory
	var memType, title, source, createdAt sql.NullString
	var confidenceScore sql.NullFloat64
	var sensitivityLevel sql.NullInt64
	if err := r.Scan(&m.MemoryID, &m.UserID, &memType, &title, &m.Content, &source, &confidenceScore, &sensitivityLevel, &createdAt); err != nil {
		return nil, err
	}
	if memType.Valid {
		m.MemoryType = &memType.String
	}
	if title.Valid {
		m.Title = &title.String
	}
	if source.Valid {
		m.Source = &source.String
	}
	if confidenceScore.Valid {
		m.ConfidenceScore = &confidenceScore.Float64
	}
	if sensitivityLevel.Valid {
		v := int(sensitivityLevel.Int64)
		m.SensitivityLevel = &v
	}
	if createdAt.Valid {
		m.CreatedAt = &createdAt.String
	}
	return &m, nil
}
