package interestcategory

type InterestCategory struct {
	CategoryID    string  `json:"category_id"`
	ParentID      *string `json:"parent_id,omitempty"`
	Name          string  `json:"name"`
	Slug          *string `json:"slug,omitempty"`
	CategoryGroup *string `json:"category_group,omitempty"`
	IsSensitive   *bool   `json:"is_sensitive,omitempty"`
}
