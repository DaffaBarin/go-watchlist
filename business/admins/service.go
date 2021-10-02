package admins

import (
	"time"

	"go-watchlist/app/middlewares"
	"go-watchlist/business"
	"go-watchlist/helper/encrypt"
)

type AdminService struct {
	repository     Repository
	contextTimeout time.Duration
	jwtAuth        *middlewares.ConfigJWT
}

func NewAdminService(repo Repository, timeout time.Duration, jwtauth *middlewares.ConfigJWT) Service {
	return &AdminService{
		repository:     repo,
		contextTimeout: timeout,
		jwtAuth:        jwtauth,
	}
}

// Business logic for register and login
func (servAdmin *AdminService) Login(username, password string) (Domain, error) {
	if username == "" {
		return Domain{}, business.ErrUsernameNotFound
	}
	if password == "" {
		return Domain{}, business.ErrPasswordNotFound
	}

	admin, err := servAdmin.repository.Login(username, password)
	if err != nil {
		return Domain{}, err
	}
	validPass := encrypt.CheckPasswordHash(password, admin.Password)
	if !validPass {
		return Domain{}, business.ErrWrongPassword
	}
	if err != nil {
		return Domain{}, err
	}
	admin.Token = servAdmin.jwtAuth.GenerateToken(admin.ID, "admin")
	if err != nil {
		return Domain{}, err
	}
	return admin, nil
}

func (servAdmin *AdminService) Register(domain *Domain) (Domain, error) {
	if domain.Username == "" {
		return Domain{}, business.ErrUsernameNotFound
	}
	if domain.Email == "" {
		return Domain{}, business.ErrEmailNotFound
	}
	if domain.Password == "" {
		return Domain{}, business.ErrPasswordNotFound
	}
	encryptedPass, _ := encrypt.HashPassword(domain.Password)
	domain.Password = encryptedPass
	admin, err := servAdmin.repository.Register(domain)

	if err != nil {
		return Domain{}, err
	}
	return admin, nil
}
