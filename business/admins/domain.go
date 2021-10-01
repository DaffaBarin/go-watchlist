package admins

import (
	"context"
	"time"
)

type Domain struct {
	ID        int
	Username  string
	Email     string
	Password  string
	Token     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Service interface {
	Register(ctx context.Context, domain Domain) (Domain, error)
	Login(ctx context.Context, email string, password string) (Domain, error)
}

type Repository interface {
	Register(ctx context.Context, domain Domain) (Domain, error)
	Login(ctx context.Context, email string, password string) (Domain, error)
}