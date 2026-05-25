package userprofile

import "context"

type Repository interface {
	List(ctx context.Context) ([]UserProfile, error)
	ListByUserID(ctx context.Context, userID string) ([]UserProfile, error)
	GetByID(ctx context.Context, id string) (*UserProfile, error)
	Create(ctx context.Context, p *UserProfile) error
	Update(ctx context.Context, p *UserProfile) error
	Delete(ctx context.Context, id string) error
}
