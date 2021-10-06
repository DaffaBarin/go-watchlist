package products

import (
	"time"
)

type Domain struct {
	ID              int
	Name            string
	Activation_Time int
	Price           int
	Description     string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type Service interface {
	CreateProduct(domain *Domain) (Domain, error)
	Delete(id int) (Domain, error)
	GetAllProduct() ([]Domain, error)
	GetProductByID(id int) (Domain, error)
	Update(id int, domain *Domain) (Domain, error)
}

type Repository interface {
	CreateProduct(domain *Domain) (Domain, error)
	Delete(id int) (Domain, error)
	GetAllProduct() ([]Domain, error)
	GetProductByID(id int) (Domain, error)
	Update(id int, domain *Domain) (Domain, error)
}
