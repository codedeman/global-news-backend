package userprofile

type UserProfile struct {
	ProfileID      string  `json:"profile_id"`
	UserID         string  `json:"user_id"`
	AgeRange       *string `json:"age_range,omitempty"`
	GenderOptional *string `json:"gender_optional,omitempty"`
	Country        *string `json:"country,omitempty"`
	Persona        *string `json:"persona,omitempty"`
	LifeStage      *string `json:"life_stage,omitempty"`
	UpdatedAt      *string `json:"updated_at,omitempty"`
}
