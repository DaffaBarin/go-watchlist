package users

import (
	"time"
)

type Domain struct {
	ID             int
	Username       string
	First_Name     string
	Last_Name      string
	Email          string
	Password       string
	DOB            string
	Gender         string
	Premium_Status bool
	Token          string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type Service interface {
	Register(domain *Domain) (Domain, error)
	Login(username string, password string) (Domain, error)
}

type Repository interface {
	Register(domain *Domain) (Domain, error)
	Login(username string, password string) (Domain, error)
}
