package users

import (
	"go-watchlist/business/users"
	"time"
)

type Users struct {
	ID             int    `gorm:"primaryKey"`
	Username       string `gorm:"unique"`
	First_Name     string
	Last_Name      string
	Email          string `gorm:"unique"`
	Password       string
	DOB            string
	Gender         string
	Premium_Status bool
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func toDomain(user Users) users.Domain {
	return users.Domain{
		ID:             user.ID,
		Username:       user.Username,
		First_Name:     user.First_Name,
		Last_Name:      user.Last_Name,
		Email:          user.Email,
		Password:       user.Password,
		DOB:            user.DOB,
		Gender:         user.Gender,
		Premium_Status: user.Premium_Status,
		CreatedAt:      user.CreatedAt,
		UpdatedAt:      user.UpdatedAt,
	}
}

func fromDomain(domain users.Domain) Users {
	return Users{
		ID:             domain.ID,
		Username:       domain.Username,
		First_Name:     domain.First_Name,
		Last_Name:      domain.Last_Name,
		Email:          domain.Email,
		Password:       domain.Password,
		DOB:            domain.DOB,
		Gender:         domain.Gender,
		Premium_Status: domain.Premium_Status,
		CreatedAt:      domain.CreatedAt,
		UpdatedAt:      domain.UpdatedAt,
	}
}
