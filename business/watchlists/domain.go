package watchlists

import (
	"time"
)

type Domain struct {
	ID              int
	UserID          int
	Name            string
	Unwatched_media []MediaStruct
	Watched_media   []MediaStruct
	Overview        string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type MediaStruct struct {
	ID       int
	Type     string
	Name     string
	Overview string
}

type Service interface {
	Create(userID int, domain *Domain) (Domain, error)
	GetAllByUserID(userID int) ([]Domain, error)
	GetByID(userID, id int) (Domain, error)
	InsertMedia(watchlistId int, mediaID int) (Domain, error)
	UpdateMedia(userID int, watchlistID int, mediaID int) (Domain, error)
	// DeleteWatchlist(userID int, watchlistID int)
}

type Repository interface {
	Create(userID int, domain *Domain) (Domain, error)
	GetAllByUserID(userID int) ([]Domain, error)
	GetByID(userID, watchlistId int) (Domain, error)
	InsertMedia(watchlistId int, mediaID int) (Domain, error)
	UpdateMedia(userID int, watchlistID int, mediaID int) (Domain, error)
	// DeleteWatchlist(userID int, watchlistID int)
}
