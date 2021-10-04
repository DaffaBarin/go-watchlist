package medias

import (
	"time"

	"go-watchlist/business"
)

type MediaService struct {
	repository     Repository
	contextTimeout time.Duration
}

func NewMediaService(repo Repository, timeout time.Duration) Service {
	return &MediaService{
		repository:     repo,
		contextTimeout: timeout,
	}
}

// Business logic for medias crud
func (servAdmin *MediaService) Create(domain *Domain) (Domain, error) {

	media, err := servAdmin.repository.Create(domain)
	if err != nil {
		return Domain{}, business.ErrInternalServer
	}

	return media, nil
}

func (servAdmin *MediaService) GetAll() ([]Domain, error) {

	media, _ := servAdmin.repository.GetAll()
	if media == nil {
		return nil, business.ErrNotFound
	}

	return media, nil
}

func (servAdmin *MediaService) GetByID(id int) (Domain, error) {

	media, err := servAdmin.repository.GetByID(id)
	if err != nil {
		return Domain{}, business.ErrNotFound
	}

	return media, nil
}
