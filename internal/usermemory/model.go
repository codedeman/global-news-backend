package usermemory

type UserMemory struct {
	MemoryID         string   `json:"memory_id"`
	UserID           string   `json:"user_id"`
	MemoryType       *string  `json:"memory_type,omitempty"`
	Title            *string  `json:"title,omitempty"`
	Content          string   `json:"content"`
	Source           *string  `json:"source,omitempty"`
	ConfidenceScore  *float64 `json:"confidence_score,omitempty"`
	SensitivityLevel *int     `json:"sensitivity_level,omitempty"`
	CreatedAt        *string  `json:"created_at,omitempty"`
}
