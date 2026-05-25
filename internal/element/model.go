package element

import "encoding/json"

type Element struct {
	ID               string           `json:"id"`
	ScreenID         string           `json:"screen_id"`
	ParentID         *string          `json:"parent_id,omitempty"`
	Name             string           `json:"name"`
	Type             string           `json:"type"`
	Selector         string           `json:"selector"`
	TextLabel        *string          `json:"text_label,omitempty"`
	Position         *json.RawMessage `json:"position,omitempty"`
	VisibleCondition *json.RawMessage `json:"visible_condition,omitempty"`
}
