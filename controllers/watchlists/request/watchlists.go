package request

import (
	"go-watchlist/business/watchlists"
)

type CreateWatchlist struct {
	Name     string `json:"name"`
	Overview string `json:"overview"`
}

type InsertMedia struct {
	MediaID int    `json:"media_id"`
	Type    string `json:"media_type"`
}

type UpdateMedia struct {
	MediaId int `json:"media_id"`
}

func (req *CreateWatchlist) ToDomain() *watchlists.Domain {
	return &watchlists.Domain{
		Name:     req.Name,
		Overview: req.Overview,
	}
}
