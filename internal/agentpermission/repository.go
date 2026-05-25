package agentpermission

import "context"

type Repository interface {
	List(ctx context.Context) ([]AgentPermission, error)
	ListByUserID(ctx context.Context, userID string) ([]AgentPermission, error)
	GetByID(ctx context.Context, id string) (*AgentPermission, error)
	Create(ctx context.Context, p *AgentPermission) error
	Update(ctx context.Context, p *AgentPermission) error
	Delete(ctx context.Context, id string) error
}
