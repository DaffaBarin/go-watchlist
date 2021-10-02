package drivers

import (
	adminDomain "go-watchlist/business/admins"
	adminDB "go-watchlist/drivers/databases/admins"

	userDomain "go-watchlist/business/users"
	userDB "go-watchlist/drivers/databases/users"

	"gorm.io/gorm"
)

func NewAdminRepository(conn *gorm.DB) adminDomain.Repository {
	return adminDB.NewMysqlAdminRepository(conn)
}

func NewUserRepository(conn *gorm.DB) userDomain.Repository {
	return userDB.NewMysqlUserRepository(conn)
}
