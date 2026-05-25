package userbehavior

type UserBehaviorSignal struct {
	SignalID          string   `json:"signal_id"`
	UserID            string   `json:"user_id"`
	SignalType        *string  `json:"signal_type,omitempty"`
	RawEventRef       *string  `json:"raw_event_ref,omitempty"`
	RelatedCategoryID *string  `json:"related_category_id,omitempty"`
	Intensity         *float64 `json:"intensity,omitempty"`
	Frequency         *int     `json:"frequency,omitempty"`
	ConfidenceScore   *float64 `json:"confidence_score,omitempty"`
}
