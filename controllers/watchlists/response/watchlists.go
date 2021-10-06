package response

import (
	"go-watchlist/business/watchlists"
)

type CreateWatchlistResponse struct {
	Message string `json:"message"`
	Name    string `json:"name:"`
}

func FromDomainCreateWatchlistResponse(domain watchlists.Domain) CreateWatchlistResponse {
	return CreateWatchlistResponse{
		Message: "Watchlist Created",
		Name:    domain.Name,
	}
}

type InsertWatchlistResponse struct {
	Message string `json:"message"`
}

func FromDomainInsertWatchlistResponse(domain watchlists.Domain) CreateWatchlistResponse {
	return CreateWatchlistResponse{
		Message: "Media Inserted to watchlist",
	}
}

type UpdateMediaResponse struct {
	Message string `json:"message"`
}

func FromDomainUpdateMediaResponse(domain watchlists.Domain) CreateWatchlistResponse {
	return CreateWatchlistResponse{
		Message: "Media status changed",
	}
}
