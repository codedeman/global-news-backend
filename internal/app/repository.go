package app

import "context"

type Repository interface {
	List(ctx context.Context) ([]App, error)
	GetByID(ctx context.Context, id string) (*App, error)
	Create(ctx context.Context, a *App) error
	Update(ctx context.Context, a *App) error
	Delete(ctx context.Context, id string) error
}
