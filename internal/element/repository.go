package element

import "context"

type Repository interface {
	ListByScreen(ctx context.Context, screenID string) ([]Element, error)
	GetByID(ctx context.Context, id string) (*Element, error)
	Create(ctx context.Context, e *Element) error
	Update(ctx context.Context, e *Element) error
	Delete(ctx context.Context, id string) error
}
