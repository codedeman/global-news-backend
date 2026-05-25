package action

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

func (s *mysqlStore) List(ctx context.Context) ([]Action, error) {
	rows, err := s.db.QueryContext(ctx,
		"SELECT id, name, type, target_element_id, parameters, preconditions, expected_outcome, priority FROM actions",
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var actions []Action
	for rows.Next() {
		a, err := scanAction(rows)
		if err != nil {
			return nil, err
		}
		actions = append(actions, *a)
	}
	return actions, rows.Err()
}

func (s *mysqlStore) GetByID(ctx context.Context, id string) (*Action, error) {
	rows, err := s.db.QueryContext(ctx,
		"SELECT id, name, type, target_element_id, parameters, preconditions, expected_outcome, priority FROM actions WHERE id = ?",
		id,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, nil
	}
	return scanAction(rows)
}

func (s *mysqlStore) Create(ctx context.Context, a *Action) error {
	p, pre, eo := marshalJSON(a.Parameters), marshalJSON(a.Preconditions), marshalJSON(a.ExpectedOutcome)
	_, err := s.db.ExecContext(ctx,
		"INSERT INTO actions (id, name, type, target_element_id, parameters, preconditions, expected_outcome, priority) VALUES (?,?,?,?,?,?,?,?)",
		a.ID, a.Name, a.Type, a.TargetElementID, p, pre, eo, a.Priority,
	)
	return err
}

func (s *mysqlStore) Update(ctx context.Context, a *Action) error {
	p, pre, eo := marshalJSON(a.Parameters), marshalJSON(a.Preconditions), marshalJSON(a.ExpectedOutcome)
	_, err := s.db.ExecContext(ctx,
		"UPDATE actions SET name=?, type=?, target_element_id=?, parameters=?, preconditions=?, expected_outcome=?, priority=? WHERE id=?",
		a.Name, a.Type, a.TargetElementID, p, pre, eo, a.Priority, a.ID,
	)
	return err
}

func (s *mysqlStore) Delete(ctx context.Context, id string) error {
	_, err := s.db.ExecContext(ctx, "DELETE FROM actions WHERE id = ?", id)
	return err
}

type scanner interface {
	Scan(dest ...any) error
}

func scanAction(r scanner) (*Action, error) {
	var a Action
	var params, pre, eo sql.NullString
	if err := r.Scan(&a.ID, &a.Name, &a.Type, &a.TargetElementID, &params, &pre, &eo, &a.Priority); err != nil {
		return nil, err
	}
	if params.Valid {
		raw := json.RawMessage(params.String)
		a.Parameters = &raw
	}
	if pre.Valid {
		raw := json.RawMessage(pre.String)
		a.Preconditions = &raw
	}
	if eo.Valid {
		raw := json.RawMessage(eo.String)
		a.ExpectedOutcome = &raw
	}
	return &a, nil
}

func marshalJSON(v *json.RawMessage) *string {
	if v == nil {
		return nil
	}
	s := string(*v)
	return &s
}
