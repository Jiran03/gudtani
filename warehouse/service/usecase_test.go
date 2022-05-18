package service_test

import (
	"errors"
	"testing"

	"github.com/Jiran03/gudhani/warehouse/domain"
	"github.com/Jiran03/gudhani/warehouse/domain/mocks"
	"github.com/Jiran03/gudhani/warehouse/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	warehouseRepo    mocks.Repository
	warehouseService domain.Service
	warehouseDomain  domain.Warehouse
)

// func TestMain(m *testing.M) {
// 	warehouseService = service.NewWarehouseService(&warehouseRepo)
// 	warehouseDomain = domain.Warehouse{
// 		Id:          1,
// 		UserId:      1,
// 		WarehouseName: "buah Gudang",
// 	}
// }

func TestInsertData(t *testing.T) {
	warehouseService = service.NewWarehouseService(&warehouseRepo)
	warehouseDomain = domain.Warehouse{
		Id:            1,
		UserId:        1,
		WarehouseName: "Gudang Jaya Abadi",
		Capacity:      12,
		RentalPrice:   8000,
		Address:       "Palu",
	}

	t.Run("InsertData | Valid", func(t *testing.T) {
		warehouseRepo.On("Create", mock.AnythingOfType("domain.Warehouse")).Return(warehouseDomain, nil).Once()
		warehouseObj, err := warehouseService.InsertData(warehouseDomain)

		assert.Nil(t, err)
		assert.Contains(t, warehouseObj.WarehouseName, "Gudang")
	})

	t.Run("InsertData | Invalid", func(t *testing.T) {
		warehouseRepo.On("Create", mock.AnythingOfType("domain.Warehouse")).Return(domain.Warehouse{}, errors.New("error")).Once()
		_, err := warehouseService.InsertData(warehouseDomain)

		assert.Error(t, err)
	})
}

func TestGetAllData(t *testing.T) {
	warehouseService = service.NewWarehouseService(&warehouseRepo)
	warehouseDomains := []domain.Warehouse{
		{
			Id:            1,
			UserId:        1,
			WarehouseName: "Gudang Jaya Abadi",
			Capacity:      12,
			RentalPrice:   8000,
			Address:       "Palu",
		},
	}

	t.Run("GetAllData | Valid", func(t *testing.T) {
		warehouseRepo.On("Get").Return(warehouseDomains, nil).Once()
		warehouseObj, err := warehouseService.GetAllData()

		assert.Nil(t, err)
		assert.Equal(t, 1, len(warehouseObj))
	})

	t.Run("GetAllData | Invalid", func(t *testing.T) {
		warehouseRepo.On("Get").Return([]domain.Warehouse{}, errors.New("error")).Once()
		_, err := warehouseService.GetAllData()

		assert.Error(t, err)
	})
}

func TestGetDataByID(t *testing.T) {
	warehouseService = service.NewWarehouseService(&warehouseRepo)
	warehouseDomain = domain.Warehouse{
		Id:            1,
		UserId:        1,
		WarehouseName: "Gudang Jaya Abadi",
		Capacity:      12,
		RentalPrice:   8000,
		Address:       "Palu",
	}

	t.Run("GetDataByID | Valid", func(t *testing.T) {
		warehouseRepo.On("GetByID", mock.AnythingOfType("int")).Return(warehouseDomain, nil).Once()
		warehouseObj, err := warehouseService.GetDataByID(warehouseDomain.Id)

		assert.Nil(t, err)
		assert.Contains(t, warehouseObj.WarehouseName, "Gudang")
	})

	t.Run("GetDataByID | Invalid", func(t *testing.T) {
		warehouseRepo.On("GetByID", mock.AnythingOfType("int")).Return(domain.Warehouse{}, errors.New("error")).Once()
		_, err := warehouseService.GetDataByID(warehouseDomain.Id)

		assert.Error(t, err)
	})
}

func TestGetDataByAddress(t *testing.T) {
	warehouseService = service.NewWarehouseService(&warehouseRepo)
	warehouseDomains := []domain.Warehouse{
		{
			Id:            1,
			UserId:        1,
			WarehouseName: "Gudang Jaya Abadi",
			Capacity:      12,
			RentalPrice:   8000,
			Address:       "Palu",
		},
	}

	t.Run("GetDataByID | Valid", func(t *testing.T) {
		warehouseRepo.On("GetByAddress", mock.AnythingOfType("string")).Return(warehouseDomains, nil).Once()
		warehouseObj, err := warehouseService.GetDataByAddress(warehouseDomains[0].Address)

		assert.Nil(t, err)
		assert.Equal(t, 1, len(warehouseObj))
	})

	t.Run("GetDataByID | Invalid", func(t *testing.T) {
		warehouseRepo.On("GetByAddress", mock.AnythingOfType("string")).Return([]domain.Warehouse{}, errors.New("error")).Once()
		_, err := warehouseService.GetDataByAddress(warehouseDomains[0].Address)

		assert.Error(t, err)
	})
}

func TestUpdateData(t *testing.T) {
	warehouseService = service.NewWarehouseService(&warehouseRepo)
	warehouseDomain = domain.Warehouse{
		Id:            1,
		UserId:        1,
		WarehouseName: "Gudang Jaya Abadi",
		Capacity:      12,
		RentalPrice:   8000,
		Address:       "Palu",
	}

	t.Run("UpdateData | Valid", func(t *testing.T) {
		warehouseRepo.On("GetByID", mock.AnythingOfType("int")).Return(warehouseDomain, nil).Once()
		warehouseRepo.On("Update", mock.AnythingOfType("int"), warehouseDomain).Return(warehouseDomain, nil).Once()
		warehouseObj, err := warehouseService.UpdateData(warehouseDomain.Id, warehouseDomain)

		assert.Nil(t, err)
		assert.Contains(t, warehouseObj.WarehouseName, "Gudang")
	})

	t.Run("UpdateData | Invalid", func(t *testing.T) {
		warehouseRepo.On("GetByID", mock.AnythingOfType("int")).Return(warehouseDomain, nil).Once()
		warehouseRepo.On("Update", mock.AnythingOfType("int"), warehouseDomain).Return(domain.Warehouse{}, errors.New("error")).Once()
		_, err := warehouseService.UpdateData(warehouseDomain.Id, warehouseDomain)

		assert.Error(t, err)
	})
}

func TestDeleteData(t *testing.T) {
	warehouseService = service.NewWarehouseService(&warehouseRepo)
	warehouseDomain = domain.Warehouse{
		Id:            1,
		UserId:        1,
		WarehouseName: "Gudang Jaya Abadi",
		Capacity:      12,
		RentalPrice:   8000,
		Address:       "Palu",
	}

	t.Run("DeleteData | Valid", func(t *testing.T) {
		warehouseRepo.On("Delete", mock.AnythingOfType("int")).Return(nil).Once()
		err := warehouseService.DeleteData(warehouseDomain.Id)

		assert.Nil(t, err)
	})

	t.Run("DeleteData | Invalid", func(t *testing.T) {
		warehouseRepo.On("Delete", mock.AnythingOfType("int")).Return(errors.New("error")).Once()
		err := warehouseService.DeleteData(warehouseDomain.Id)

		assert.Error(t, err)
	})
}
