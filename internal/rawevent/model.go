package rawevent

import "encoding/json"

type RawEvent struct {
	EventID     string           `json:"event_id"`
	UserID      string           `json:"user_id"`
	EventType   string           `json:"event_type"`
	Source      *string          `json:"source,omitempty"`
	PayloadJSON *json.RawMessage `json:"payload_json,omitempty"`
	OccurredAt  *string          `json:"occurred_at,omitempty"`
	CreatedAt   *string          `json:"created_at,omitempty"`
}
