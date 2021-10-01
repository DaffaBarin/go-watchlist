package admins

import (
	"context"
	"errors"
	"time"
)

type AdminService struct {
	repository           Repository
	contextTimeout time.Duration
}

func NewAdminService(repo Repository, timeout time.Duration) Service {
	return &AdminService{
		repository:     repo,
		contextTimeout: timeout,
	}
}

// Business logic for register and login
func (servAdmin *AdminService) Login(ctx context.Context, email string, password string) (Domain, error) {
	if email == "" {
		return Domain{}, errors.New("username empty")
	}
	if password == "" {
		return Domain{}, errors.New("password empty")
	}

	admin, err := servAdmin.repository.Login(ctx, email, password)

	if err != nil {
		return Domain{}, err
	}
	// admin.Token, err = servAdmin.ConfigJWT.GenerateToken(admin.ID, "admin")
	if err != nil {
		return Domain{}, err
	}
	return admin, nil
}

func (servAdmin *AdminService) Register(ctx context.Context, domain Domain) (Domain, error) {
	if domain.Username == "" {
		return Domain{}, errors.New("username empty")
	}
	if domain.Email == "" {
		return Domain{}, errors.New("email empty")
	}
	if domain.Password == "" {
		return Domain{}, errors.New("password empty")
	}
	admin, err := servAdmin.repository.Register(ctx, domain)

	if err != nil {
		return Domain{}, err
	}
	return admin, nil
}