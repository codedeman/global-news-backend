package agentprofile

import "context"

type Repository interface {
	List(ctx context.Context) ([]AgentProfile, error)
	GetByID(ctx context.Context, id string) (*AgentProfile, error)
	Create(ctx context.Context, p *AgentProfile) error
	Update(ctx context.Context, p *AgentProfile) error
	Delete(ctx context.Context, id string) error
}
