package action

import "context"

type Repository interface {
	List(ctx context.Context) ([]Action, error)
	GetByID(ctx context.Context, id string) (*Action, error)
	Create(ctx context.Context, a *Action) error
	Update(ctx context.Context, a *Action) error
	Delete(ctx context.Context, id string) error
}
