package agentpermission

import "encoding/json"

type AgentPermission struct {
	PermissionID     string           `json:"permission_id"`
	UserID           *string          `json:"user_id,omitempty"`
	DataScope        *string          `json:"data_scope,omitempty"`
	PermissionLevel  *int             `json:"permission_level,omitempty"`
	SensitivityLimit *int             `json:"sensitivity_limit,omitempty"`
	AllowedAgentTypes *json.RawMessage `json:"allowed_agent_types,omitempty"`
}
