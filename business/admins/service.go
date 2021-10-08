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
		return Domain{}, business.ErrEmptyForm
	}
	if password == "" {
		return Domain{}, business.ErrEmptyForm
	}

	admin, err := servAdmin.repository.Login(username, password)
	if err != nil {
		return Domain{}, business.ErrUser
	}
	validPass := encrypt.CheckPasswordHash(password, admin.Password)
	if !validPass {
		return Domain{}, business.ErrUser
	}
	admin.Token = servAdmin.jwtAuth.GenerateToken(admin.ID, "admin")
	return admin, nil
}

func (servAdmin *AdminService) Register(domain *Domain) (Domain, error) {
	if domain.Username == "" {
		return Domain{}, business.ErrEmptyForm
	}
	if domain.Email == "" {
		return Domain{}, business.ErrEmptyForm
	}
	if domain.Password == "" {
		return Domain{}, business.ErrEmptyForm
	}
	encryptedPass, _ := encrypt.HashPassword(domain.Password)
	domain.Password = encryptedPass
	admin, err := servAdmin.repository.Register(domain)

	if err != nil {
		return Domain{}, err
	}
	return admin, nil
}
