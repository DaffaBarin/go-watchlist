package drivers

import (
	adminDomain "go-watchlist/business/admins"
	adminDB "go-watchlist/drivers/databases/admins"

	"gorm.io/gorm"
)

func NewAdminRepository(conn *gorm.DB) adminDomain.Repository {
	return adminDB.NewMysqlAdminRepository(conn)
}
