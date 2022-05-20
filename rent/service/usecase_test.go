package service_test

import (
	"errors"
	"testing"

	"github.com/Jiran03/gudtani/rent/domain"
	"github.com/Jiran03/gudtani/rent/domain/mocks"
	"github.com/Jiran03/gudtani/rent/service"
	wareDomain "github.com/Jiran03/gudtani/warehouse/domain"
	wareMock "github.com/Jiran03/gudtani/warehouse/domain/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	rentRepo         mocks.Repository
	warehouseService wareMock.Service

	rentService     domain.Service
	rentDomain      domain.Rent
	warehouseDomain wareDomain.Warehouse
)

func TestInsertData(t *testing.T) {
	rentService = service.NewRentService(&rentRepo, &warehouseService)
	rentDomain = domain.Rent{
		ID:          1,
		ProductID:   1,
		WarehouseID: 2,
		Weight:      12,
		Period:      14,
		Status:      "sedang dalam masa sewa",
		TotalPrice:  140000,
	}

	warehouseDomain = wareDomain.Warehouse{
		Id:            1,
		UserId:        1,
		WarehouseName: "Gudang Jaya Abadi",
		Capacity:      12,
		RentalPrice:   8000,
		Address:       "Palu",
	}

	t.Run("InsertData | Valid", func(t *testing.T) {
		warehouseService.On("GetDataByID", mock.AnythingOfType("int")).Return(warehouseDomain, nil).Once()
		warehouseService.On("UpdateDataCapacity", mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(nil)
		rentRepo.On("Create", mock.AnythingOfType("domain.Rent")).Return(rentDomain, nil).Once()
		rentObj, err := rentService.InsertData(rentDomain)

		assert.Nil(t, err)
		assert.Contains(t, rentObj.Status, "sewa")
	})

	t.Run("InsertData | Invalid", func(t *testing.T) {
		warehouseService.On("GetDataByID", mock.AnythingOfType("int")).Return(wareDomain.Warehouse{}, errors.New("error")).Once()
		warehouseService.On("UpdateDataCapacity", mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(errors.New("error"))
		rentRepo.On("Create", mock.AnythingOfType("domain.Rent")).Return(domain.Rent{}, errors.New("error")).Once()
		_, err := rentService.InsertData(rentDomain)

		assert.Error(t, err)
	})
}

func TestGetAllData(t *testing.T) {
	rentService = service.NewRentService(&rentRepo, &warehouseService)
	rentDomains := []domain.Rent{
		{
			ID:          1,
			ProductID:   1,
			WarehouseID: 1,
			Weight:      12,
			Period:      14,
			Status:      "sedang dalam masa sewa",
			TotalPrice:  140000,
		},
	}

	t.Run("GetAllData | Valid", func(t *testing.T) {
		rentRepo.On("Get").Return(rentDomains, nil).Once()
		rentObj, err := rentService.GetAllData()

		assert.Nil(t, err)
		assert.Equal(t, 1, len(rentObj))
	})

	t.Run("GetAllData | Invalid", func(t *testing.T) {
		rentRepo.On("Get").Return([]domain.Rent{}, errors.New("error")).Once()
		_, err := rentService.GetAllData()

		assert.Error(t, err)
	})
}

func TestGetDataByID(t *testing.T) {
	rentService = service.NewRentService(&rentRepo, &warehouseService)
	rentDomain = domain.Rent{
		ID:          1,
		ProductID:   1,
		WarehouseID: 2,
		Weight:      12,
		Period:      14,
		Status:      "sedang dalam masa sewa",
		TotalPrice:  140000,
	}
	t.Run("GetDataByID | Valid", func(t *testing.T) {
		rentRepo.On("GetByID", mock.AnythingOfType("int")).Return(rentDomain, nil).Once()
		rentObj, err := rentService.GetDataByID(rentDomain.ID)

		assert.Nil(t, err)
		assert.Contains(t, rentObj.Status, "sewa")
	})

	t.Run("GetDataByID | Invalid", func(t *testing.T) {
		rentRepo.On("GetByID", mock.AnythingOfType("int")).Return(domain.Rent{}, errors.New("error")).Once()
		_, err := rentService.GetDataByID(rentDomain.ID)

		assert.Error(t, err)
	})
}

func TestUpdateData(t *testing.T) {
	rentService = service.NewRentService(&rentRepo, &warehouseService)
	rentDomain = domain.Rent{
		ID:          1,
		ProductID:   1,
		WarehouseID: 2,
		Weight:      12,
		Period:      14,
		Status:      "sedang dalam masa sewa",
		TotalPrice:  140000,
	}

	warehouseDomain = wareDomain.Warehouse{
		Id:            1,
		UserId:        1,
		WarehouseName: "Gudang Jaya Abadi",
		Capacity:      12,
		RentalPrice:   8000,
		Address:       "Palu",
	}

	t.Run("UpdateData | Valid", func(t *testing.T) {
		rentRepo.On("GetByID", mock.AnythingOfType("int")).Return(rentDomain, nil).Once()
		warehouseService.On("GetDataByID", mock.AnythingOfType("int")).Return(warehouseDomain, nil).Once()
		warehouseService.On("UpdateDataCapacity", mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(nil)
		rentRepo.On("Update", mock.AnythingOfType("int"), mock.AnythingOfType("domain.Rent")).Return(rentDomain, nil).Once()
		rentObj, err := rentService.UpdateData(rentDomain.ID, rentDomain)

		assert.Nil(t, err)
		assert.Contains(t, rentObj.Status, "sewa")
	})

	t.Run("UpdateData | Invalid", func(t *testing.T) {
		rentRepo.On("GetByID", mock.AnythingOfType("int")).Return(domain.Rent{}, errors.New("error")).Once()
		warehouseService.On("GetDataByID", mock.AnythingOfType("int")).Return(wareDomain.Warehouse{}, errors.New("error")).Once()
		warehouseService.On("UpdateDataCapacity", mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(errors.New("error"))
		rentRepo.On("Update", mock.AnythingOfType("int"), mock.AnythingOfType("domain.Rent")).Return(domain.Rent{}, errors.New("error")).Once()
		_, err := rentService.UpdateData(rentDomain.ID, rentDomain)

		assert.Error(t, err)
	})
}

func TestDeleteData(t *testing.T) {
	rentService = service.NewRentService(&rentRepo, &warehouseService)
	rentDomain = domain.Rent{
		ID:          1,
		ProductID:   1,
		WarehouseID: 2,
		Weight:      12,
		Period:      14,
		Status:      "sedang dalam masa sewa",
		TotalPrice:  140000,
	}

	t.Run("DeleteData | Valid", func(t *testing.T) {
		rentRepo.On("Delete", mock.AnythingOfType("int")).Return(nil).Once()
		err := rentService.DeleteData(rentDomain.ID)

		assert.Nil(t, err)
	})

	t.Run("DeleteData | Invalid", func(t *testing.T) {
		rentRepo.On("Delete", mock.AnythingOfType("int")).Return(errors.New("error")).Once()
		err := rentService.DeleteData(rentDomain.ID)

		assert.Error(t, err)
	})
}
