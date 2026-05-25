package rawevent

import (
	"context"
	"database/sql"
	"encoding/json"
)

type mysqlStore struct {
	db *sql.DB
}

func NewStore(db *sql.DB) Repository {
	return &mysqlStore{db: db}
}

const selectCols = "SELECT event_id, user_id, event_type, source, payload_json, occurred_at, created_at FROM raw_events"

func (s *mysqlStore) List(ctx context.Context) ([]RawEvent, error) {
	rows, err := s.db.QueryContext(ctx, selectCols)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanAll(rows)
}

func (s *mysqlStore) ListByUserID(ctx context.Context, userID string) ([]RawEvent, error) {
	rows, err := s.db.QueryContext(ctx, selectCols+" WHERE user_id = ?", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanAll(rows)
}

func (s *mysqlStore) GetByID(ctx context.Context, id string) (*RawEvent, error) {
	rows, err := s.db.QueryContext(ctx, selectCols+" WHERE event_id = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if !rows.Next() {
		return nil, nil
	}
	return scanEvent(rows)
}

func (s *mysqlStore) Create(ctx context.Context, e *RawEvent) error {
	_, err := s.db.ExecContext(ctx,
		"INSERT INTO raw_events (event_id, user_id, event_type, source, payload_json, occurred_at) VALUES (?, ?, ?, ?, ?, ?)",
		e.EventID, e.UserID, e.EventType, e.Source, marshalJSON(e.PayloadJSON), e.OccurredAt,
	)
	return err
}

func (s *mysqlStore) Update(ctx context.Context, e *RawEvent) error {
	_, err := s.db.ExecContext(ctx,
		"UPDATE raw_events SET user_id = ?, event_type = ?, source = ?, payload_json = ?, occurred_at = ? WHERE event_id = ?",
		e.UserID, e.EventType, e.Source, marshalJSON(e.PayloadJSON), e.OccurredAt, e.EventID,
	)
	return err
}

func (s *mysqlStore) Delete(ctx context.Context, id string) error {
	_, err := s.db.ExecContext(ctx, "DELETE FROM raw_events WHERE event_id = ?", id)
	return err
}

type scanner interface {
	Scan(dest ...any) error
}

func scanAll(rows *sql.Rows) ([]RawEvent, error) {
	var items []RawEvent
	for rows.Next() {
		e, err := scanEvent(rows)
		if err != nil {
			return nil, err
		}
		items = append(items, *e)
	}
	return items, rows.Err()
}

func scanEvent(r scanner) (*RawEvent, error) {
	var e RawEvent
	var source, payloadJSON, occurredAt, createdAt sql.NullString
	if err := r.Scan(&e.EventID, &e.UserID, &e.EventType, &source, &payloadJSON, &occurredAt, &createdAt); err != nil {
		return nil, err
	}
	if source.Valid {
		e.Source = &source.String
	}
	if payloadJSON.Valid {
		raw := json.RawMessage(payloadJSON.String)
		e.PayloadJSON = &raw
	}
	if occurredAt.Valid {
		e.OccurredAt = &occurredAt.String
	}
	if createdAt.Valid {
		e.CreatedAt = &createdAt.String
	}
	return &e, nil
}

func marshalJSON(v *json.RawMessage) *string {
	if v == nil {
		return nil
	}
	s := string(*v)
	return &s
}
