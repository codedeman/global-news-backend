package goalaction

import "context"

type Repository interface {
	ListByGoal(ctx context.Context, goalID string) ([]GoalAction, error)
	GetByID(ctx context.Context, id int64) (*GoalAction, error)
	Create(ctx context.Context, ga *GoalAction) error
	Update(ctx context.Context, ga *GoalAction) error
	Delete(ctx context.Context, id int64) error
}
