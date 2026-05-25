package userinterest

type UserInterest struct {
	UserInterestID  string   `json:"user_interest_id"`
	UserID          string   `json:"user_id"`
	CategoryID      *string  `json:"category_id,omitempty"`
	InterestStrength *float64 `json:"interest_strength,omitempty"`
	Source          *string  `json:"source,omitempty"`
	ConfidenceScore *float64 `json:"confidence_score,omitempty"`
	UserConfirmed   *string  `json:"user_confirmed,omitempty"`
}
