package request

import "go-watchlist/business/admins"

// request body for login
type AdminsLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// request body for register
type Admins struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// turn register body to domain
func (req *Admins) ToDomain() *admins.Domain {
	return &admins.Domain{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}
}
