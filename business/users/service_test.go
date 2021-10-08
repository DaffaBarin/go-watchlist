package users_test

import (
	"go-watchlist/app/middlewares"
	"go-watchlist/business"
	"go-watchlist/business/users"
	userMock "go-watchlist/business/users/mocks"
	"go-watchlist/helper/encrypt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	userMockRepo userMock.Repository
	jwtAuth      *middlewares.ConfigJWT
	userService  users.Service
)

func TestMain(m *testing.M) {
	jwtAuth := &middlewares.ConfigJWT{
		SecretJWT:       "testmock123",
		ExpiresDuration: 1,
	}
	userService = users.NewUserService(&userMockRepo, 2, jwtAuth)

	m.Run()
}

// go test -v ./...
func TestLogin(t *testing.T) {
	t.Run("Test case 1 | Valid login", func(t *testing.T) {
		encryptedPass, _ := encrypt.HashPassword("daffa123")
		// anggap si user domain database
		UserDomain := users.Domain{
			ID:        1,
			Username:  "nama123",
			Email:     "nama@email123.com",
			Password:  encryptedPass,
			Token:     "token123",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		expectedReturnService := users.Domain{
			Username: "user123",
			Password: "daffa123",
		}

		userMockRepo.On("Login", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(UserDomain, nil).Once()

		resp, err := userService.Login(expectedReturnService.Username, expectedReturnService.Password)
		assert.Nil(t, err)
		assert.NotEmpty(t, resp)
	})
	t.Run("Test case 2 | Password Invalid", func(t *testing.T) {
		encryptedPass, _ := encrypt.HashPassword("daffa123")
		// anggap si user domain database
		UserDomain := users.Domain{
			ID:        1,
			Username:  "nama123",
			Email:     "nama@email123.com",
			Password:  encryptedPass,
			Token:     "token123",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		expectedReturnService := users.Domain{
			Username: "nama123",
			Password: "daffawrong",
		}

		userMockRepo.On("Login", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(UserDomain, nil).Once()

		_, err := userService.Login(expectedReturnService.Username, expectedReturnService.Password)
		assert.Equal(t, err, business.ErrUser)
	})

	t.Run("Test case 3 | Empty Password", func(t *testing.T) {
		encryptedPass, _ := encrypt.HashPassword("daffa123")
		// anggap si user domain database
		UserDomain := users.Domain{
			ID:        1,
			Username:  "nama123",
			Email:     "nama@email123.com",
			Password:  encryptedPass,
			Token:     "token123",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		expectedReturnService := users.Domain{
			Username: "nama123",
			Password: "",
		}

		userMockRepo.On("Login", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(UserDomain, nil).Once()

		_, err := userService.Login(expectedReturnService.Username, expectedReturnService.Password)
		assert.Equal(t, err, business.ErrEmptyForm)
	})
	t.Run("Test case 4 | Empty Username", func(t *testing.T) {
		encryptedPass, _ := encrypt.HashPassword("daffa123")
		// anggap si user domain database
		UserDomain := users.Domain{
			ID:        1,
			Username:  "nama123",
			Email:     "nama@email123.com",
			Password:  encryptedPass,
			Token:     "token123",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		expectedReturnService := users.Domain{
			Username: "",
			Password: encryptedPass,
		}

		userMockRepo.On("Login", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(UserDomain, nil).Once()

		_, err := userService.Login(expectedReturnService.Username, expectedReturnService.Password)
		assert.Equal(t, err, business.ErrEmptyForm)
	})

	t.Run("Test case 4 | Invalid Req", func(t *testing.T) {
		// encryptedPass, _ := encrypt.HashPassword("daffa123")
		// anggap si user domain database
		UserDomain := users.Domain{}

		expectedReturnService := users.Domain{
			Username: "nama",
			Password: "pass123",
		}

		userMockRepo.On("Login", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(UserDomain, business.ErrUser).Once()

		_, err := userService.Login(expectedReturnService.Username, expectedReturnService.Password)
		assert.NotNil(t, err)
		assert.NotEqual(t, UserDomain.Username, expectedReturnService.Username)
	})
}

func TestRegister(t *testing.T) {
	t.Run("Test case 1 | Valid Register", func(t *testing.T) {
		encryptedPass, _ := encrypt.HashPassword("daffa123")
		// anggap si user domain database
		UserDomain := users.Domain{
			ID:        1,
			Username:  "nama123",
			Email:     "nama@email123.com",
			Password:  encryptedPass,
			Token:     "token123",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		expectedReturnService := users.Domain{
			Username: "user123",
			Email:    "nama@email123.com",
			Password: encryptedPass,
		}

		userMockRepo.On("Register", mock.Anything).Return(UserDomain, nil).Once()

		resp, err := userService.Register(&expectedReturnService)
		assert.Nil(t, err)
		assert.NotEmpty(t, resp)

	})

	t.Run("Test case 2 | Empty Username", func(t *testing.T) {
		encryptedPass, _ := encrypt.HashPassword("daffa123")
		// anggap si user domain database
		UserDomain := users.Domain{
			ID:        1,
			Username:  "nama123",
			Email:     "nama@email123.com",
			Password:  encryptedPass,
			Token:     "token123",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		expectedReturnService := users.Domain{
			Username: "",
			Email:    "nama@email123.com",
			Password: encryptedPass,
		}

		userMockRepo.On("Register", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(UserDomain, nil).Once()

		_, err := userService.Register(&expectedReturnService)
		assert.Equal(t, err, business.ErrEmptyForm)
		// err == business.ErrPassword == true; ok pass

	})

	t.Run("Test case 3 | Empty Email", func(t *testing.T) {
		encryptedPass, _ := encrypt.HashPassword("daffa123")
		// anggap si user domain database
		UserDomain := users.Domain{
			ID:        1,
			Username:  "nama123",
			Email:     "nama@email123.com",
			Password:  encryptedPass,
			Token:     "token123",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		expectedReturnService := users.Domain{
			Username: "nama123",
			Email:    "",
			Password: encryptedPass,
		}

		userMockRepo.On("Register", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(UserDomain, nil).Once()

		_, err := userService.Register(&expectedReturnService)
		assert.Equal(t, err, business.ErrEmptyForm)
	})

	t.Run("Test case 4 | Empty Password", func(t *testing.T) {
		encryptedPass, _ := encrypt.HashPassword("daffa123")
		// anggap si user domain database
		UserDomain := users.Domain{
			ID:        1,
			Username:  "nama123",
			Email:     "nama@email123.com",
			Password:  encryptedPass,
			Token:     "token123",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		expectedReturnService := users.Domain{
			Username: "nama123",
			Email:    "nama@email123.com",
			Password: "",
		}

		userMockRepo.On("Register", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(UserDomain, nil).Once()

		_, err := userService.Register(&expectedReturnService)
		assert.Equal(t, err, business.ErrEmptyForm)
	})
}
