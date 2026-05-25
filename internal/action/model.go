package action

import "encoding/json"

type Action struct {
	ID              string           `json:"id"`
	Name            string           `json:"name"`
	Type            string           `json:"type"`
	TargetElementID *string          `json:"target_element_id,omitempty"`
	Parameters      *json.RawMessage `json:"parameters,omitempty"`
	Preconditions   *json.RawMessage `json:"preconditions,omitempty"`
	ExpectedOutcome *json.RawMessage `json:"expected_outcome,omitempty"`
	Priority        int              `json:"priority"`
}
