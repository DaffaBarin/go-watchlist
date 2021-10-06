package watchlists

import (
	"time"

	"go-watchlist/business"
	"go-watchlist/business/medias"
	"go-watchlist/business/users"
)

type WatchlistService struct {
	repository      Repository
	userRepository  users.Repository
	mediaRepository medias.Repository
	contextTimeout  time.Duration
}

func NewWatchlistService(repo Repository, userRepo users.Repository, medaiRepo medias.Repository, timeout time.Duration) Service {
	return &WatchlistService{
		repository:      repo,
		userRepository:  userRepo,
		mediaRepository: medaiRepo,
		contextTimeout:  timeout,
	}
}

// Business logic for medias crud
func (servWatchlist *WatchlistService) Create(userID int, domain *Domain) (Domain, error) {

	media, err := servWatchlist.repository.Create(userID, domain)
	if err != nil {
		return Domain{}, business.ErrInternalServer
	}

	return media, nil
}

func (servWatchlist *WatchlistService) GetAllByUserID(userID int) ([]Domain, error) {

	media, _ := servWatchlist.repository.GetAllByUserID(userID)
	if media == nil {
		return nil, business.ErrNotFound
	}

	return media, nil
}

func (servWatchlist *WatchlistService) GetByID(userID, id int) (Domain, error) {

	media, err := servWatchlist.repository.GetByID(userID, id)
	if err != nil {
		return Domain{}, business.ErrNotFound
	}

	return media, nil
}

func (servWatchlist *WatchlistService) InsertMedia(watchlistId, mediaID int) (Domain, error) {

	media, err := servWatchlist.repository.InsertMedia(watchlistId, mediaID)
	if err != nil {
		return Domain{}, business.ErrInternalServer
	}

	return media, nil
}

func (servWatchlist *WatchlistService) UpdateMedia(userID int, watchlistID int, mediaID int) (Domain, error) {

	media, err := servWatchlist.repository.UpdateMedia(userID, watchlistID, mediaID)
	if err != nil {
		return Domain{}, business.ErrInternalServer
	}

	return media, nil
}
