package user

type User struct {
	UserID      string  `json:"user_id"`
	Name        string  `json:"name"`
	Email       *string `json:"email,omitempty"`
	Language    *string `json:"language,omitempty"`
	Timezone    *string `json:"timezone,omitempty"`
	PrivacyMode *string `json:"privacy_mode,omitempty"`
	CreatedAt   *string `json:"created_at,omitempty"`
}
