package service_test

import (
	"errors"
	"testing"

	authMiddleware "github.com/Jiran03/gudhani/auth/middleware"
	"github.com/Jiran03/gudhani/user/domain"
	"github.com/Jiran03/gudhani/user/domain/mocks"
	"github.com/Jiran03/gudhani/user/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	userRepo    mocks.Repository
	userService domain.Service
	userDomain  domain.User
)

func TestInsertData(t *testing.T) {
	userService = service.NewUserService(&userRepo, authMiddleware.ConfigJWT{})
	userDomain = domain.User{
		ID:       1,
		Name:     "fulan",
		Email:    "fulan@mail.com",
		Password: "fulanpass",
		Address:  "Palu",
		Gender:   "L",
		Role:     "Petani, Pemilik Gudang",
	}

	t.Run("InsertData | Valid", func(t *testing.T) {
		userRepo.On("Create", mock.AnythingOfType("domain.User")).Return(userDomain, nil).Once()
		userObj, err := userService.InsertData(userDomain)

		assert.Nil(t, err)
		assert.Contains(t, userObj.Name, "fulan")
	})

	t.Run("InsertData | Invalid", func(t *testing.T) {
		userRepo.On("Create", mock.AnythingOfType("domain.User")).Return(domain.User{}, errors.New("error")).Once()
		_, err := userService.InsertData(userDomain)

		assert.Error(t, err)
	})
}

func TestGetAllData(t *testing.T) {
	userService = service.NewUserService(&userRepo, authMiddleware.ConfigJWT{})
	userDomains := []domain.User{
		{
			ID:       1,
			Name:     "fulan",
			Email:    "fulan@mail.com",
			Password: "fulanpass",
			Address:  "Palu",
			Gender:   "L",
			Role:     "Petani, Pemilik Gudang",
		},
	}

	t.Run("GetAllData | Valid", func(t *testing.T) {
		userRepo.On("Get").Return(userDomains, nil).Once()
		userObj, err := userService.GetAllData()

		assert.Nil(t, err)
		assert.Equal(t, 1, len(userObj))
	})

	t.Run("GetAllData | Invalid", func(t *testing.T) {
		userRepo.On("Get").Return([]domain.User{}, errors.New("error")).Once()
		_, err := userService.GetAllData()

		assert.Error(t, err)
	})
}

func TestGetDataByID(t *testing.T) {
	userService = service.NewUserService(&userRepo, authMiddleware.ConfigJWT{})
	userDomain = domain.User{
		ID:       1,
		Name:     "fulan",
		Email:    "fulan@mail.com",
		Password: "fulanpass",
		Address:  "Palu",
		Gender:   "L",
		Role:     "Petani, Pemilik Gudang",
	}

	t.Run("GetDataByID | Valid", func(t *testing.T) {
		userRepo.On("GetByID", mock.AnythingOfType("int")).Return(userDomain, nil).Once()
		userObj, err := userService.GetByID(userDomain.ID)

		assert.Nil(t, err)
		assert.Contains(t, userObj.Name, "fulan")
	})

	t.Run("GetDataByID | Invalid", func(t *testing.T) {
		userRepo.On("GetByID", mock.AnythingOfType("int")).Return(domain.User{}, errors.New("error")).Once()
		_, err := userService.GetByID(userDomain.ID)

		assert.Error(t, err)
	})
}

func TestGetByEmailPassword(t *testing.T) {
	userService = service.NewUserService(&userRepo, authMiddleware.ConfigJWT{})
	userDomain = domain.User{
		ID:       1,
		Name:     "fulan",
		Email:    "fulan@mail.com",
		Password: "fulanpass",
		Address:  "Palu",
		Gender:   "L",
		Role:     "Petani, Pemilik Gudang",
	}

	t.Run("GetDataByEmailPassword | Valid", func(t *testing.T) {
		userRepo.On("GetByEmailPassword", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(userDomain, nil).Once()
		_, userRole, err := userService.GetByEmailPassword(userDomain.Email, userDomain.Password)

		assert.Nil(t, err)
		// assert.Contains(t, userID, userDomain.ID)
		assert.Contains(t, userRole, "Petani")
	})

	t.Run("GetDataByEmailPassword | Invalid", func(t *testing.T) {
		userRepo.On("GetByEmailPassword", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(domain.User{}, errors.New("error")).Once()
		_, _, err := userService.GetByEmailPassword(userDomain.Email, userDomain.Password)

		assert.Error(t, err)
	})
}

func TestUpdateData(t *testing.T) {
	userService = service.NewUserService(&userRepo, authMiddleware.ConfigJWT{})
	userDomain = domain.User{
		ID:       1,
		Name:     "fulan",
		Email:    "fulan@mail.com",
		Password: "fulanpass",
		Address:  "Palu",
		Gender:   "L",
		Role:     "Petani, Pemilik Gudang",
	}

	t.Run("UpdateData | Valid", func(t *testing.T) {
		userRepo.On("GetByID", mock.AnythingOfType("int")).Return(userDomain, nil).Once()
		userRepo.On("Update", mock.AnythingOfType("int"), userDomain).Return(userDomain, nil).Once()
		userObj, err := userService.UpdateData(userDomain.ID, userDomain)

		assert.Nil(t, err)
		assert.Contains(t, userObj.Name, "fulan")
	})

	t.Run("UpdateData | Invalid", func(t *testing.T) {
		userRepo.On("GetByID", mock.AnythingOfType("int")).Return(userDomain, nil).Once()
		userRepo.On("Update", mock.AnythingOfType("int"), userDomain).Return(domain.User{}, errors.New("error")).Once()
		_, err := userService.UpdateData(userDomain.ID, userDomain)

		assert.Error(t, err)
	})
}

func TestDeleteData(t *testing.T) {
	userService = service.NewUserService(&userRepo, authMiddleware.ConfigJWT{})
	userDomain = domain.User{
		ID:       1,
		Name:     "fulan",
		Email:    "fulan@mail.com",
		Password: "fulanpass",
		Address:  "Palu",
		Gender:   "L",
		Role:     "Petani, Pemilik Gudang",
	}

	t.Run("DeleteData | Valid", func(t *testing.T) {
		userRepo.On("Delete", mock.AnythingOfType("int")).Return(nil).Once()
		err := userService.DeleteData(userDomain.ID)

		assert.Nil(t, err)
	})

	t.Run("DeleteData | Invalid", func(t *testing.T) {
		userRepo.On("Delete", mock.AnythingOfType("int")).Return(errors.New("error")).Once()
		err := userService.DeleteData(userDomain.ID)

		assert.Error(t, err)
	})
}
