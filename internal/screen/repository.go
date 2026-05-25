package screen

import "context"

type Repository interface {
	List(ctx context.Context) ([]Screen, error)
	ListByApp(ctx context.Context, appID string) ([]Screen, error)
	GetByID(ctx context.Context, id string) (*Screen, error)
	Create(ctx context.Context, s *Screen) error
	Update(ctx context.Context, s *Screen) error
	Delete(ctx context.Context, id string) error
}
