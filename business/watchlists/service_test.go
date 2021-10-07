package watchlists_test

import (
	// "go-watchlist/app/middlewares"
	// "go-watchlist/business"

	mediaMock "go-watchlist/business/medias/mocks"
	userMock "go-watchlist/business/users/mocks"
	"go-watchlist/business/watchlists"
	watchlistMock "go-watchlist/business/watchlists/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	watchlistMockRepo watchlistMock.Repository
	userMockRepo      userMock.Repository
	mediaMockrepo     mediaMock.Repository
	watchlistService  watchlists.Service
)

func TestMain(m *testing.M) {
	watchlistService = watchlists.NewWatchlistService(&watchlistMockRepo, &userMockRepo, &mediaMockrepo, 2)

	m.Run()
}

func TestCreate(t *testing.T) {
	t.Run("Test case 1 | Succes Create Watchlist", func(t *testing.T) {
		WatchlistDomain := watchlists.Domain{
			ID:        1,
			UserID:    1,
			Name:      "Sebuah watchlist",
			Overview:  "Watchlist untuk nonton malming",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		expectedReturnService := watchlists.Domain{
			ID:        1,
			UserID:    1,
			Name:      "Sebuah watchlist",
			Overview:  "Watchlist untuk nonton malming",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		watchlistMockRepo.On("Create", mock.AnythingOfType("int"), mock.Anything).Return(WatchlistDomain, nil).Once()
		resp, err := watchlistService.Create(expectedReturnService.ID, &expectedReturnService)
		assert.Nil(t, err)
		assert.NotEmpty(t, resp)
	})
}

func TestGetAllByUserID(t *testing.T) {
	t.Run("Test case 1 | Succes Get All By Used ID", func(t *testing.T) {
		WatchlistDomain := []watchlists.Domain{
			{ID: 1,
				UserID:    1,
				Name:      "Sebuah watchlist",
				Overview:  "Watchlist untuk nonton malming",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now()},
		}
		// expectedReturnService := []watchlists.Domain{
		// {	ID:        1,
		// 	UserID:    1,
		// 	Name:      "Sebuah watchlist",
		// 	Overview:  "Watchlist untuk nonton malming",
		// 	CreatedAt: time.Now(),
		// 	UpdatedAt: time.Now()},
		// }
		watchlistMockRepo.On("GetAllByUserID", mock.AnythingOfType("int"), mock.Anything).Return(WatchlistDomain, nil).Once()
		resp, err := watchlistService.GetAllByUserID(1)
		assert.Nil(t, err)
		assert.NotEmpty(t, resp)
	})
}

func TestGetByID(t *testing.T) {
	t.Run("Test case 1 | Succes Get All By User ID", func(t *testing.T) {
		WatchlistDomain := watchlists.Domain{
			ID:        1,
			UserID:    1,
			Name:      "Sebuah watchlist",
			Overview:  "Watchlist untuk nonton malming",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		watchlistMockRepo.On("GetByID", mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(WatchlistDomain, nil).Once()
		resp, err := watchlistService.GetByID(1, 1)
		assert.NotNil(t, resp)
		assert.Nil(t, err)
	})

}

func TestInsertMedia(t *testing.T) {
	t.Run("Test case 1 | Succes Insert Media", func(t *testing.T) {
		WatchlistDomain := watchlists.Domain{
			ID:        1,
			UserID:    1,
			Name:      "Sebuah watchlist",
			Overview:  "Watchlist untuk nonton malming",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		watchlistMockRepo.On("InsertMedia", mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(WatchlistDomain, nil).Once()
		resp, err := watchlistService.InsertMedia(1, 1)
		assert.Nil(t, err)
		assert.NotEmpty(t, resp)
	})
}

func TestUpdateMedia(t *testing.T) {
	t.Run("Test case 1 | Succes Update Media", func(t *testing.T) {
		WatchlistDomain := watchlists.Domain{
			ID:        1,
			UserID:    1,
			Name:      "Sebuah watchlist",
			Overview:  "Watchlist untuk nonton malming",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		watchlistMockRepo.On("UpdateMedia", mock.AnythingOfType("int"), mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(WatchlistDomain, nil).Once()
		resp, err := watchlistService.UpdateMedia(1, 1, 1)
		assert.NotNil(t, resp)
		assert.Nil(t, err)
	})
}
