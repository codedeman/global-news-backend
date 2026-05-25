package johari

import "context"

type Repository interface {
	List(ctx context.Context) ([]JohariUnderstanding, error)
	ListByUserID(ctx context.Context, userID string) ([]JohariUnderstanding, error)
	GetByID(ctx context.Context, id string) (*JohariUnderstanding, error)
	Create(ctx context.Context, j *JohariUnderstanding) error
	Update(ctx context.Context, j *JohariUnderstanding) error
	Delete(ctx context.Context, id string) error
}
