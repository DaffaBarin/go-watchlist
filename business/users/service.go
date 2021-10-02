package users

import (
	"time"

	"go-watchlist/app/middlewares"
	"go-watchlist/business"
	"go-watchlist/helper/encrypt"
)

type UserService struct {
	repository     Repository
	contextTimeout time.Duration
	jwtAuth        *middlewares.ConfigJWT
}

func NewUserService(repo Repository, timeout time.Duration, jwtauth *middlewares.ConfigJWT) Service {
	return &UserService{
		repository:     repo,
		contextTimeout: timeout,
		jwtAuth:        jwtauth,
	}
}

// Business logic for register and login
func (servUser *UserService) Login(username, password string) (Domain, error) {
	if username == "" {
		return Domain{}, business.ErrUsernameNotFound
	}
	if password == "" {
		return Domain{}, business.ErrPasswordNotFound
	}

	user, err := servUser.repository.Login(username, password)
	if err != nil {
		return Domain{}, err
	}
	validPass := encrypt.CheckPasswordHash(password, user.Password)
	if !validPass {
		return Domain{}, business.ErrWrongPassword
	}
	if err != nil {
		return Domain{}, err
	}
	user.Token = servUser.jwtAuth.GenerateToken(user.ID, "user")
	if err != nil {
		return Domain{}, err
	}
	return user, nil
}

func (servUser *UserService) Register(domain *Domain) (Domain, error) {
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
	user, err := servUser.repository.Register(domain)

	if err != nil {
		return Domain{}, err
	}
	return user, nil
}
