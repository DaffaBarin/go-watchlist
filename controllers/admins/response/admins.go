package response

import (
	"go-watchlist/business/admins"
	"time"
)

type AdminRegisterResponse struct {
	Message   string    `json:"message"`
	ID        int       `json:"id:"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FromDomainRegister(domain admins.Domain) AdminRegisterResponse {
	return AdminRegisterResponse{
		Message:   "Registration Success",
		ID:        domain.ID,
		Username:  domain.Username,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

type AdminLoginResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

func FromDomainLogin(domain admins.Domain) AdminLoginResponse {
	return AdminLoginResponse{
		Message: "Login Success",
		Token:   domain.Token,
	}
}
