package watchlists

import (
	"go-watchlist/business/watchlists"
	"go-watchlist/drivers/databases/medias"
	"time"
)

type Watchlists struct {
	ID              int `gorm:"primary_key"`
	Name            string
	UserID          int
	Unwatched_media []medias.Medias          `gorm:"many2many:unwatched_pivots"`
	Watched_media   []medias.Medias          `gorm:"many2many:watched_pivots"`
	Unwatched       []watchlists.MediaStruct `gorm:"-"`
	Watched         []watchlists.MediaStruct `gorm:"-"`
	Overview        string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type WatchedPivot struct {
	Watchlists_id int `gorm:"watchlists_id"`
	Medias_id     int `gorm:"medias_id"`
}
type UnwatchedPivot struct {
	Watchlists_id int `gorm:"watchlists_id"`
	Medias_id     int `gorm:"medias_id"`
}

func toDomain(watchlist Watchlists) watchlists.Domain {
	return watchlists.Domain{
		ID:              watchlist.ID,
		UserID:          watchlist.UserID,
		Name:            watchlist.Name,
		Watched_media:   watchlist.Watched,
		Unwatched_media: watchlist.Unwatched,
		Overview:        watchlist.Overview,
		CreatedAt:       watchlist.CreatedAt,
		UpdatedAt:       watchlist.UpdatedAt,
	}
}

func toListDomain(domain []Watchlists) (result []watchlists.Domain) {
	result = []watchlists.Domain{}
	for _, domain := range domain {
		result = append(result, toDomain(domain))
	}
	return result
}

func fromDomain(domain watchlists.Domain) Watchlists {
	return Watchlists{
		ID:        domain.ID,
		UserID:    domain.UserID,
		Name:      domain.Name,
		Overview:  domain.Overview,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
