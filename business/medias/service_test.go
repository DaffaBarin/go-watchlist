package medias_test

import (
	"go-watchlist/business/medias"
	mediaMock "go-watchlist/business/medias/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	mediaMockRepo mediaMock.Repository
	mediaService  medias.Service
)

func TestMain(m *testing.M) {
	mediaService = medias.NewMediaService(&mediaMockRepo, 2)

	m.Run()
}

func TestCreate(t *testing.T) {
	t.Run("Test case 1 | Succes Create Media", func(t *testing.T) {
		MediaDomain := medias.Domain{
			ID:   1,
			Type: "Movie",
		}

		expectedReturnService := medias.Domain{
			ID:   1,
			Type: "Movie",
		}

		mediaMockRepo.On("Create", mock.Anything).Return(MediaDomain, nil).Once()
		resp, err := mediaService.Create(&expectedReturnService)
		assert.Nil(t, err)
		assert.NotEmpty(t, resp)
	})
}

func TestGetAll(t *testing.T) {
	t.Run("Test case 1 | Succes Get All", func(t *testing.T) {
		MediaDomain := []medias.Domain{
			{ID: 1,
				Type: "Movie"},
		}

		mediaMockRepo.On("GetAll", mock.Anything).Return(MediaDomain, nil).Once()
		resp, err := mediaService.GetAll()
		assert.Nil(t, err)
		assert.NotEmpty(t, resp)
	})

	t.Run("Test case 2 | Media empty", func(t *testing.T) {
		MediaDomain := []medias.Domain{}

		mediaMockRepo.On("GetAll", mock.Anything).Return(MediaDomain, nil).Once()
		resp, _ := mediaService.GetAll()
		assert.Equal(t, len(resp), 0)
	})
}

func TestGetByID(t *testing.T) {
	t.Run("Test case 1 | Succes Get By ID", func(t *testing.T) {
		MediaDomain := medias.Domain{
			ID:   1,
			Type: "Movie",
		}
		expectedReturnService := medias.Domain{
			ID:   1,
			Type: "Movie",
		}
		mediaMockRepo.On("GetByID", mock.Anything).Return(MediaDomain, nil).Once()
		resp, err := mediaService.GetByID(expectedReturnService.ID)
		assert.Nil(t, err)
		assert.NotEmpty(t, resp)
	})
}
