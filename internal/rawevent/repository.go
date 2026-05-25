package rawevent

import "context"

type Repository interface {
	List(ctx context.Context) ([]RawEvent, error)
	ListByUserID(ctx context.Context, userID string) ([]RawEvent, error)
	GetByID(ctx context.Context, id string) (*RawEvent, error)
	Create(ctx context.Context, e *RawEvent) error
	Update(ctx context.Context, e *RawEvent) error
	Delete(ctx context.Context, id string) error
}
