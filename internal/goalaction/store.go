package goalaction

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

func (s *mysqlStore) ListByGoal(ctx context.Context, goalID string) ([]GoalAction, error) {
	rows, err := s.db.QueryContext(ctx,
		"SELECT id, goal_id, action_id, step_order, context_condition FROM goal_action_mapping WHERE goal_id = ? ORDER BY step_order",
		goalID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []GoalAction
	for rows.Next() {
		ga, err := scanGoalAction(rows)
		if err != nil {
			return nil, err
		}
		items = append(items, *ga)
	}
	return items, rows.Err()
}

func (s *mysqlStore) GetByID(ctx context.Context, id int64) (*GoalAction, error) {
	rows, err := s.db.QueryContext(ctx,
		"SELECT id, goal_id, action_id, step_order, context_condition FROM goal_action_mapping WHERE id = ?",
		id,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, nil
	}
	return scanGoalAction(rows)
}

func (s *mysqlStore) Create(ctx context.Context, ga *GoalAction) error {
	cc := marshalJSON(ga.ContextCondition)
	result, err := s.db.ExecContext(ctx,
		"INSERT INTO goal_action_mapping (goal_id, action_id, step_order, context_condition) VALUES (?,?,?,?)",
		ga.GoalID, ga.ActionID, ga.StepOrder, cc,
	)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	ga.ID = id
	return nil
}

func (s *mysqlStore) Update(ctx context.Context, ga *GoalAction) error {
	cc := marshalJSON(ga.ContextCondition)
	_, err := s.db.ExecContext(ctx,
		"UPDATE goal_action_mapping SET goal_id=?, action_id=?, step_order=?, context_condition=? WHERE id=?",
		ga.GoalID, ga.ActionID, ga.StepOrder, cc, ga.ID,
	)
	return err
}

func (s *mysqlStore) Delete(ctx context.Context, id int64) error {
	_, err := s.db.ExecContext(ctx, "DELETE FROM goal_action_mapping WHERE id = ?", id)
	return err
}

type scanner interface {
	Scan(dest ...any) error
}

func scanGoalAction(r scanner) (*GoalAction, error) {
	var ga GoalAction
	var cc sql.NullString
	if err := r.Scan(&ga.ID, &ga.GoalID, &ga.ActionID, &ga.StepOrder, &cc); err != nil {
		return nil, err
	}
	if cc.Valid {
		raw := json.RawMessage(cc.String)
		ga.ContextCondition = &raw
	}
	return &ga, nil
}

func marshalJSON(v *json.RawMessage) *string {
	if v == nil {
		return nil
	}
	s := string(*v)
	return &s
}
