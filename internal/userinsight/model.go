package userinsight

type UserInsight struct {
	InsightID         string  `json:"insight_id"`
	UserID            string  `json:"user_id"`
	InsightType       *string `json:"insight_type,omitempty"`
	Title             *string `json:"title,omitempty"`
	Description       *string `json:"description,omitempty"`
	JohariZone        *string `json:"johari_zone,omitempty"`
	Status            *string `json:"status,omitempty"`
	RecommendedAction *string `json:"recommended_action,omitempty"`
}
