package agentpermission

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

const selectCols = "SELECT permission_id, user_id, data_scope, permission_level, sensitivity_limit, allowed_agent_types FROM agent_permissions"

func (s *mysqlStore) List(ctx context.Context) ([]AgentPermission, error) {
	rows, err := s.db.QueryContext(ctx, selectCols)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanAll(rows)
}

func (s *mysqlStore) ListByUserID(ctx context.Context, userID string) ([]AgentPermission, error) {
	rows, err := s.db.QueryContext(ctx, selectCols+" WHERE user_id = ?", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanAll(rows)
}

func (s *mysqlStore) GetByID(ctx context.Context, id string) (*AgentPermission, error) {
	rows, err := s.db.QueryContext(ctx, selectCols+" WHERE permission_id = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if !rows.Next() {
		return nil, nil
	}
	return scanPermission(rows)
}

func (s *mysqlStore) Create(ctx context.Context, p *AgentPermission) error {
	_, err := s.db.ExecContext(ctx,
		"INSERT INTO agent_permissions (permission_id, user_id, data_scope, permission_level, sensitivity_limit, allowed_agent_types) VALUES (?, ?, ?, ?, ?, ?)",
		p.PermissionID, p.UserID, p.DataScope, p.PermissionLevel, p.SensitivityLimit, marshalJSON(p.AllowedAgentTypes),
	)
	return err
}

func (s *mysqlStore) Update(ctx context.Context, p *AgentPermission) error {
	_, err := s.db.ExecContext(ctx,
		"UPDATE agent_permissions SET user_id = ?, data_scope = ?, permission_level = ?, sensitivity_limit = ?, allowed_agent_types = ? WHERE permission_id = ?",
		p.UserID, p.DataScope, p.PermissionLevel, p.SensitivityLimit, marshalJSON(p.AllowedAgentTypes), p.PermissionID,
	)
	return err
}

func (s *mysqlStore) Delete(ctx context.Context, id string) error {
	_, err := s.db.ExecContext(ctx, "DELETE FROM agent_permissions WHERE permission_id = ?", id)
	return err
}

type scanner interface {
	Scan(dest ...any) error
}

func scanAll(rows *sql.Rows) ([]AgentPermission, error) {
	var items []AgentPermission
	for rows.Next() {
		p, err := scanPermission(rows)
		if err != nil {
			return nil, err
		}
		items = append(items, *p)
	}
	return items, rows.Err()
}

func scanPermission(r scanner) (*AgentPermission, error) {
	var p AgentPermission
	var userID, dataScope, allowedTypes sql.NullString
	var permLevel, sensitivityLimit sql.NullInt64
	if err := r.Scan(&p.PermissionID, &userID, &dataScope, &permLevel, &sensitivityLimit, &allowedTypes); err != nil {
		return nil, err
	}
	if userID.Valid {
		p.UserID = &userID.String
	}
	if dataScope.Valid {
		p.DataScope = &dataScope.String
	}
	if permLevel.Valid {
		v := int(permLevel.Int64)
		p.PermissionLevel = &v
	}
	if sensitivityLimit.Valid {
		v := int(sensitivityLimit.Int64)
		p.SensitivityLimit = &v
	}
	if allowedTypes.Valid {
		raw := json.RawMessage(allowedTypes.String)
		p.AllowedAgentTypes = &raw
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
