package request

import "go-watchlist/business/admins"

// request body for login
type AdminLogin struct {
	Username string `json:"email"`
	Password string `json:"password"`
}

// request body for register
type AdminRegister struct{
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// turn register body to domain
func fromDomain(admin AdminLogin) admins.Domain {
	return admins.Domain{
		Username:  admin.Username,
		Password:  admin.Password,
	}
}