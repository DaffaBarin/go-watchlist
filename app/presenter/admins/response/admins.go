package response

import (
	"go-watchlist/business/admins"
	"time"
)

type AdminRegisterResponse struct{
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Username  string    `json:"username"`
}

func FromDomainRegister(domain admins.Domain) AdminRegisterResponse {
	return AdminRegisterResponse{
		Message :  "Registration Success",
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		Username:  domain.Username,
	}
}

type AdminLoginResponse struct{
	Message string    `json:"message"`
	Token   string    `json:"token"`
}

func FromDomainLogin(domain admins.Domain) AdminLoginResponse {
	return AdminLoginResponse{
		Message : "Login Success",
		Token   : domain.Token,
	}
}