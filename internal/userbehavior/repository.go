package userbehavior

import "context"

type Repository interface {
	List(ctx context.Context) ([]UserBehaviorSignal, error)
	ListByUserID(ctx context.Context, userID string) ([]UserBehaviorSignal, error)
	GetByID(ctx context.Context, id string) (*UserBehaviorSignal, error)
	Create(ctx context.Context, s *UserBehaviorSignal) error
	Update(ctx context.Context, s *UserBehaviorSignal) error
	Delete(ctx context.Context, id string) error
}
