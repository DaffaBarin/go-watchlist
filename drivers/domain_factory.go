package drivers

import (
	adminDomain "go-watchlist/business/admins"
	adminDB "go-watchlist/drivers/databases/admins"

	userDomain "go-watchlist/business/users"
	userDB "go-watchlist/drivers/databases/users"

	mediaDomain "go-watchlist/business/medias"
	mediaDB "go-watchlist/drivers/databases/medias"

	watchlistDomain "go-watchlist/business/watchlists"
	watchlistDB "go-watchlist/drivers/databases/watchlists"

	"gorm.io/gorm"
)

func NewAdminRepository(conn *gorm.DB) adminDomain.Repository {
	return adminDB.NewMysqlAdminRepository(conn)
}

func NewUserRepository(conn *gorm.DB) userDomain.Repository {
	return userDB.NewMysqlUserRepository(conn)
}

func NewMediaRepository(conn *gorm.DB) mediaDomain.Repository {
	return mediaDB.NewMysqlMediaRepository(conn)
}

func NewWatchListRepository(conn *gorm.DB) watchlistDomain.Repository {
	return watchlistDB.NewMysqlWatchlistRepository(conn)
}
