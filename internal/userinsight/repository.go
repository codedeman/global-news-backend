package userinsight

import "context"

type Repository interface {
	List(ctx context.Context) ([]UserInsight, error)
	ListByUserID(ctx context.Context, userID string) ([]UserInsight, error)
	GetByID(ctx context.Context, id string) (*UserInsight, error)
	Create(ctx context.Context, i *UserInsight) error
	Update(ctx context.Context, i *UserInsight) error
	Delete(ctx context.Context, id string) error
}
