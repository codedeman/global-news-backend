package agentprofile

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

const selectCols = "SELECT agent_id, name, agent_type, description, allowed_interest_groups, default_permission_level FROM agent_profiles"

func (s *mysqlStore) List(ctx context.Context) ([]AgentProfile, error) {
	rows, err := s.db.QueryContext(ctx, selectCols)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanAll(rows)
}

func (s *mysqlStore) GetByID(ctx context.Context, id string) (*AgentProfile, error) {
	rows, err := s.db.QueryContext(ctx, selectCols+" WHERE agent_id = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if !rows.Next() {
		return nil, nil
	}
	return scanProfile(rows)
}

func (s *mysqlStore) Create(ctx context.Context, p *AgentProfile) error {
	_, err := s.db.ExecContext(ctx,
		"INSERT INTO agent_profiles (agent_id, name, agent_type, description, allowed_interest_groups, default_permission_level) VALUES (?, ?, ?, ?, ?, ?)",
		p.AgentID, p.Name, p.AgentType, p.Description, marshalJSON(p.AllowedInterestGroups), p.DefaultPermissionLevel,
	)
	return err
}

func (s *mysqlStore) Update(ctx context.Context, p *AgentProfile) error {
	_, err := s.db.ExecContext(ctx,
		"UPDATE agent_profiles SET name = ?, agent_type = ?, description = ?, allowed_interest_groups = ?, default_permission_level = ? WHERE agent_id = ?",
		p.Name, p.AgentType, p.Description, marshalJSON(p.AllowedInterestGroups), p.DefaultPermissionLevel, p.AgentID,
	)
	return err
}

func (s *mysqlStore) Delete(ctx context.Context, id string) error {
	_, err := s.db.ExecContext(ctx, "DELETE FROM agent_profiles WHERE agent_id = ?", id)
	return err
}

type scanner interface {
	Scan(dest ...any) error
}

func scanAll(rows *sql.Rows) ([]AgentProfile, error) {
	var items []AgentProfile
	for rows.Next() {
		p, err := scanProfile(rows)
		if err != nil {
			return nil, err
		}
		items = append(items, *p)
	}
	return items, rows.Err()
}

func scanProfile(r scanner) (*AgentProfile, error) {
	var p AgentProfile
	var agentType, description, allowedGroups sql.NullString
	var defaultLevel sql.NullInt64
	if err := r.Scan(&p.AgentID, &p.Name, &agentType, &description, &allowedGroups, &defaultLevel); err != nil {
		return nil, err
	}
	if agentType.Valid {
		p.AgentType = &agentType.String
	}
	if description.Valid {
		p.Description = &description.String
	}
	if allowedGroups.Valid {
		raw := json.RawMessage(allowedGroups.String)
		p.AllowedInterestGroups = &raw
	}
	if defaultLevel.Valid {
		v := int(defaultLevel.Int64)
		p.DefaultPermissionLevel = &v
	}
	return &p, nil
}

func marshalJSON(v *json.RawMessage) *string {
	if v == nil {
		return nil
	}
	s := string(*v)
	return &s
}
