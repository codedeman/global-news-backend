package interestcategory

import "context"

type Repository interface {
	List(ctx context.Context) ([]InterestCategory, error)
	GetByID(ctx context.Context, id string) (*InterestCategory, error)
	Create(ctx context.Context, c *InterestCategory) error
	Update(ctx context.Context, c *InterestCategory) error
	Delete(ctx context.Context, id string) error
}
