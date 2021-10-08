package admins_test

import (
	"go-watchlist/app/middlewares"
	"go-watchlist/business"
	"go-watchlist/business/admins"
	adminMock "go-watchlist/business/admins/mocks"
	"go-watchlist/helper/encrypt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	adminMockRepo adminMock.Repository
	jwtAuth       *middlewares.ConfigJWT
	adminService  admins.Service
)

func TestMain(m *testing.M) {
	jwtAuth = &middlewares.ConfigJWT{
		SecretJWT:       "testmock123",
		ExpiresDuration: 1,
	}
	adminService = admins.NewAdminService(&adminMockRepo, 2, jwtAuth)

	m.Run()
}

// go test -v ./...
func TestLogin(t *testing.T) {
	t.Run("Test case 1 | Valid login", func(t *testing.T) {
		encryptedPass, _ := encrypt.HashPassword("daffa123")
		// anggap si user domain database
		UserDomain := admins.Domain{
			ID:        1,
			Username:  "nama123",
			Email:     "nama@email123.com",
			Password:  encryptedPass,
			Token:     "token123",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		expectedReturnService := admins.Domain{
			Username: "nama123",
			Password: "daffa123",
		}

		adminMockRepo.On("Login", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(UserDomain, nil).Once()

		resp, err := adminService.Login(expectedReturnService.Username, expectedReturnService.Password)
		assert.Nil(t, err)
		// if t != err => nil
		assert.NotEmpty(t, resp)
		// iresp kosong, maka ok.

	})

	t.Run("Test case 2 | Password Invalid", func(t *testing.T) {
		encryptedPass, _ := encrypt.HashPassword("daffa123")
		// anggap si user domain database
		UserDomain := admins.Domain{
			ID:        1,
			Username:  "nama123",
			Email:     "nama@email123.com",
			Password:  encryptedPass,
			Token:     "token123",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		expectedReturnService := admins.Domain{
			Username: "user123",
			Password: "daffa234213",
		}

		adminMockRepo.On("Login", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(UserDomain, nil).Once()

		_, err := adminService.Login(expectedReturnService.Username, expectedReturnService.Password)
		assert.Equal(t, err, business.ErrUser)

	})

	t.Run("Test case 3 | Username Empty", func(t *testing.T) {
		encryptedPass, _ := encrypt.HashPassword("daffa123")
		// anggap si user domain database
		UserDomain := admins.Domain{
			ID:        1,
			Username:  "nama123",
			Email:     "nama@email123.com",
			Password:  encryptedPass,
			Token:     "token123",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		expectedReturnService := admins.Domain{
			Username: "",
			Password: encryptedPass,
		}

		adminMockRepo.On("Login", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(UserDomain, nil).Once()

		_, err := adminService.Login(expectedReturnService.Username, expectedReturnService.Password)
		assert.Equal(t, err, business.ErrEmptyForm)
	})

	t.Run("Test case 4 | Password Empty", func(t *testing.T) {
		encryptedPass, _ := encrypt.HashPassword("daffa123")
		// anggap si user domain database
		UserDomain := admins.Domain{
			ID:        1,
			Username:  "nama123",
			Email:     "nama@email123.com",
			Password:  encryptedPass,
			Token:     "token123",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		expectedReturnService := admins.Domain{
			Username: "nama123",
			Password: "",
		}

		adminMockRepo.On("Login", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(UserDomain, nil).Once()

		_, err := adminService.Login(expectedReturnService.Username, expectedReturnService.Password)
		assert.Equal(t, err, business.ErrEmptyForm)
	})

	t.Run("Test case 5 | Not Found", func(t *testing.T) {
		encryptedPass, _ := encrypt.HashPassword("daffa123")
		// anggap si user domain database
		UserDomain := admins.Domain{
			ID:        1,
			Username:  "nama123",
			Email:     "nama@email123.com",
			Password:  encryptedPass,
			Token:     "token123",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		expectedReturnService := admins.Domain{
			Username: "asdasdas",
			Password: "sdasdasdasd",
		}

		adminMockRepo.On("Login", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(UserDomain, nil).Once()

		_, err := adminService.Login(expectedReturnService.Username, expectedReturnService.Password)
		assert.NotNil(t, err)
		assert.Equal(t, err, business.ErrUser)
	})
}

func TestRegister(t *testing.T) {
	t.Run("Test case 1 | Valid Register", func(t *testing.T) {
		encryptedPass, _ := encrypt.HashPassword("daffa123")
		// anggap si user domain database
		UserDomain := admins.Domain{
			ID:        1,
			Username:  "nama123",
			Email:     "nama@email123.com",
			Password:  encryptedPass,
			Token:     "token123",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		expectedReturnService := admins.Domain{
			Username: "user123",
			Email:    "nama@email123.com",
			Password: encryptedPass,
		}

		adminMockRepo.On("Register", mock.Anything).Return(UserDomain, nil).Once()

		resp, err := adminService.Register(&expectedReturnService)
		assert.Nil(t, err)
		// if t != err => nil
		assert.NotEmpty(t, resp)
		// iresp kosong, maka ok.

	})

	t.Run("Test case 2 | Empty Username", func(t *testing.T) {
		encryptedPass, _ := encrypt.HashPassword("daffa123")
		// anggap si user domain database
		UserDomain := admins.Domain{
			ID:        1,
			Username:  "nama123",
			Email:     "nama@email123.com",
			Password:  encryptedPass,
			Token:     "token123",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		expectedReturnService := admins.Domain{
			Username: "",
			Email:    "nama@email123.com",
			Password: encryptedPass,
		}

		adminMockRepo.On("Register", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(UserDomain, nil).Once()

		_, err := adminService.Register(&expectedReturnService)
		assert.Equal(t, err, business.ErrEmptyForm)
		// err == business.ErrPassword == true; ok pass

	})

	t.Run("Test case 3 | Empty Email", func(t *testing.T) {
		encryptedPass, _ := encrypt.HashPassword("daffa123")
		// anggap si user domain database
		UserDomain := admins.Domain{
			ID:        1,
			Username:  "nama123",
			Email:     "nama@email123.com",
			Password:  encryptedPass,
			Token:     "token123",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		expectedReturnService := admins.Domain{
			Username: "nama123",
			Email:    "",
			Password: encryptedPass,
		}

		adminMockRepo.On("Register", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(UserDomain, nil).Once()

		_, err := adminService.Register(&expectedReturnService)
		assert.Equal(t, err, business.ErrEmptyForm)
	})

	t.Run("Test case 4 | Empty Password", func(t *testing.T) {
		encryptedPass, _ := encrypt.HashPassword("daffa123")
		// anggap si user domain database
		UserDomain := admins.Domain{
			ID:        1,
			Username:  "nama123",
			Email:     "nama@email123.com",
			Password:  encryptedPass,
			Token:     "token123",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		expectedReturnService := admins.Domain{
			Username: "nama123",
			Email:    "nama@email123.com",
			Password: "",
		}

		adminMockRepo.On("Register", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(UserDomain, nil).Once()

		_, err := adminService.Register(&expectedReturnService)
		assert.Equal(t, err, business.ErrEmptyForm)
	})
}
