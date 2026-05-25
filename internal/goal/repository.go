package goal

import "context"

type Repository interface {
	List(ctx context.Context) ([]Goal, error)
	GetByID(ctx context.Context, id string) (*Goal, error)
	Create(ctx context.Context, g *Goal) error
	Update(ctx context.Context, g *Goal) error
	Delete(ctx context.Context, id string) error
}
