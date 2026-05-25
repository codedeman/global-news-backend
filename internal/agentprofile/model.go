package agentprofile

import "encoding/json"

type AgentProfile struct {
	AgentID                string           `json:"agent_id"`
	Name                   string           `json:"name"`
	AgentType              *string          `json:"agent_type,omitempty"`
	Description            *string          `json:"description,omitempty"`
	AllowedInterestGroups  *json.RawMessage `json:"allowed_interest_groups,omitempty"`
	DefaultPermissionLevel *int             `json:"default_permission_level,omitempty"`
}
