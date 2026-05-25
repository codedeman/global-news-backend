package goalaction

import "encoding/json"

type GoalAction struct {
	ID               int64            `json:"id"`
	GoalID           string           `json:"goal_id"`
	ActionID         string           `json:"action_id"`
	StepOrder        int              `json:"step_order"`
	ContextCondition *json.RawMessage `json:"context_condition,omitempty"`
}
