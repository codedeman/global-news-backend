package johari

type JohariUnderstanding struct {
	ID              string   `json:"id"`
	UserID          *string  `json:"user_id,omitempty"`
	TargetType      *string  `json:"target_type,omitempty"`
	TargetID        *string  `json:"target_id,omitempty"`
	JohariZone      *string  `json:"johari_zone,omitempty"`
	KnownByUser     *bool    `json:"known_by_user,omitempty"`
	KnownByAI       *bool    `json:"known_by_ai,omitempty"`
	ConfirmedByUser *bool    `json:"confirmed_by_user,omitempty"`
	ConfidenceScore *float64 `json:"confidence_score,omitempty"`
}
