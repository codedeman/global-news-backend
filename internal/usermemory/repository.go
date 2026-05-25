package usermemory

import "context"

type Repository interface {
	List(ctx context.Context) ([]UserMemory, error)
	ListByUserID(ctx context.Context, userID string) ([]UserMemory, error)
	GetByID(ctx context.Context, id string) (*UserMemory, error)
	Create(ctx context.Context, m *UserMemory) error
	Update(ctx context.Context, m *UserMemory) error
	Delete(ctx context.Context, id string) error
}
