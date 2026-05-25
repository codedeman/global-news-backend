package userinterest

import "context"

type Repository interface {
	List(ctx context.Context) ([]UserInterest, error)
	ListByUserID(ctx context.Context, userID string) ([]UserInterest, error)
	GetByID(ctx context.Context, id string) (*UserInterest, error)
	Create(ctx context.Context, i *UserInterest) error
	Update(ctx context.Context, i *UserInterest) error
	Delete(ctx context.Context, id string) error
}
