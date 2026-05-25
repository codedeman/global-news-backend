package user

import "context"

type Repository interface {
	List(ctx context.Context) ([]User, error)
	GetByID(ctx context.Context, id string) (*User, error)
	Create(ctx context.Context, u *User) error
	Update(ctx context.Context, u *User) error
	Delete(ctx context.Context, id string) error
}
