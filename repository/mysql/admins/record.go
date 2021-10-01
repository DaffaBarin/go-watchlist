package admins

import (
	"go-watchlist/business/admins"
	"time"
)

type Admins struct {
	ID int  `gorm:"primaryKey"`
	Username string `gorm:"unique"`
	Email string `gorm:"unique"`
	Password string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func toDomain(admin Admins) admins.Domain {
	return admins.Domain{
		ID:        admin.ID,
		Email:     admin.Email,
		Password:  admin.Password,
		CreatedAt: admin.CreatedAt,
		UpdatedAt: admin.UpdatedAt,
	}
}

func fromDomain(domain admins.Domain) Admins {
	return Admins{
		ID:        domain.ID,
		Email:     domain.Email,
		Password:  domain.Password,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}