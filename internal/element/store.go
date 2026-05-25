package element

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

func (s *mysqlStore) ListByScreen(ctx context.Context, screenID string) ([]Element, error) {
	rows, err := s.db.QueryContext(ctx,
		"SELECT id, screen_id, parent_id, name, type, selector, text_label, position, visible_condition FROM elements WHERE screen_id = ?",
		screenID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var elements []Element
	for rows.Next() {
		e, err := scanElement(rows)
		if err != nil {
			return nil, err
		}
		elements = append(elements, *e)
	}
	return elements, rows.Err()
}

func (s *mysqlStore) GetByID(ctx context.Context, id string) (*Element, error) {
	rows, err := s.db.QueryContext(ctx,
		"SELECT id, screen_id, parent_id, name, type, selector, text_label, position, visible_condition FROM elements WHERE id = ?",
		id,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, nil
	}
	return scanElement(rows)
}

func (s *mysqlStore) Create(ctx context.Context, e *Element) error {
	pos, vc := marshalJSON(e.Position), marshalJSON(e.VisibleCondition)
	_, err := s.db.ExecContext(ctx,
		"INSERT INTO elements (id, screen_id, parent_id, name, type, selector, text_label, position, visible_condition) VALUES (?,?,?,?,?,?,?,?,?)",
		e.ID, e.ScreenID, e.ParentID, e.Name, e.Type, e.Selector, e.TextLabel, pos, vc,
	)
	return err
}

func (s *mysqlStore) Update(ctx context.Context, e *Element) error {
	pos, vc := marshalJSON(e.Position), marshalJSON(e.VisibleCondition)
	_, err := s.db.ExecContext(ctx,
		"UPDATE elements SET screen_id=?, parent_id=?, name=?, type=?, selector=?, text_label=?, position=?, visible_condition=? WHERE id=?",
		e.ScreenID, e.ParentID, e.Name, e.Type, e.Selector, e.TextLabel, pos, vc, e.ID,
	)
	return err
}

func (s *mysqlStore) Delete(ctx context.Context, id string) error {
	_, err := s.db.ExecContext(ctx, "DELETE FROM elements WHERE id = ?", id)
	return err
}

type scanner interface {
	Scan(dest ...any) error
}

func scanElement(r scanner) (*Element, error) {
	var e Element
	var pos, vc sql.NullString
	if err := r.Scan(&e.ID, &e.ScreenID, &e.ParentID, &e.Name, &e.Type, &e.Selector, &e.TextLabel, &pos, &vc); err != nil {
		return nil, err
	}
	if pos.Valid {
		raw := json.RawMessage(pos.String)
		e.Position = &raw
	}
	if vc.Valid {
		raw := json.RawMessage(vc.String)
		e.VisibleCondition = &raw
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
