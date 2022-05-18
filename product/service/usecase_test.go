package service_test

import (
	"errors"
	"testing"

	"github.com/Jiran03/gudhani/product/domain"
	"github.com/Jiran03/gudhani/product/domain/mocks"
	"github.com/Jiran03/gudhani/product/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	productRepo    mocks.Repository
	productService domain.Service
	productDomain  domain.Product
)

// func TestMain(m *testing.M) {
// 	productService = service.NewProductService(&productRepo)
// 	productDomain = domain.Product{
// 		Id:          1,
// 		UserId:      1,
// 		ProductType: "buah anggur",
// 	}
// }

func TestInsertData(t *testing.T) {
	productService = service.NewProductService(&productRepo)
	productDomain = domain.Product{
		Id:          1,
		UserId:      1,
		ProductType: "buah anggur",
	}

	t.Run("InsertData | Valid", func(t *testing.T) {
		productRepo.On("Create", mock.AnythingOfType("domain.Product")).Return(productDomain, nil).Once()
		productObj, err := productService.InsertData(productDomain)

		assert.Nil(t, err)
		assert.Contains(t, productObj.ProductType, "anggur")
	})

	t.Run("InsertData | Invalid", func(t *testing.T) {
		productRepo.On("Create", mock.AnythingOfType("domain.Product")).Return(domain.Product{}, errors.New("error")).Once()
		_, err := productService.InsertData(productDomain)

		assert.Error(t, err)
	})
}

func TestGetAllData(t *testing.T) {
	productService = service.NewProductService(&productRepo)
	productDomains := []domain.Product{
		{
			Id:          1,
			UserId:      1,
			ProductType: "buah anggur",
		},
	}

	t.Run("GetAllData | Valid", func(t *testing.T) {
		productRepo.On("Get").Return(productDomains, nil).Once()
		productObj, err := productService.GetAllData()

		assert.Nil(t, err)
		assert.Equal(t, 1, len(productObj))
	})

	t.Run("GetAllData | Invalid", func(t *testing.T) {
		productRepo.On("Get").Return([]domain.Product{}, errors.New("error")).Once()
		_, err := productService.GetAllData()

		assert.Error(t, err)
	})
}

func TestGetDataByID(t *testing.T) {
	productService = service.NewProductService(&productRepo)
	productDomain = domain.Product{
		Id:          1,
		UserId:      1,
		ProductType: "buah anggur",
	}

	t.Run("GetDataByID | Valid", func(t *testing.T) {
		productRepo.On("GetByID", mock.AnythingOfType("int")).Return(productDomain, nil).Once()
		productObj, err := productService.GetDataByID(productDomain.Id)

		assert.Nil(t, err)
		assert.Contains(t, productObj.ProductType, "anggur")
	})

	t.Run("GetDataByID | Invalid", func(t *testing.T) {
		productRepo.On("GetByID", mock.AnythingOfType("int")).Return(domain.Product{}, errors.New("error")).Once()
		_, err := productService.GetDataByID(productDomain.Id)

		assert.Error(t, err)
	})
}

func TestUpdateData(t *testing.T) {
	productService = service.NewProductService(&productRepo)
	productDomain = domain.Product{
		Id:          1,
		UserId:      1,
		ProductType: "buah anggur",
	}

	t.Run("UpdateData | Valid", func(t *testing.T) {
		productRepo.On("GetByID", mock.AnythingOfType("int")).Return(productDomain, nil).Once()
		productRepo.On("Update", mock.AnythingOfType("int"), productDomain).Return(productDomain, nil).Once()
		productObj, err := productService.UpdateData(productDomain.Id, productDomain)

		assert.Nil(t, err)
		assert.Contains(t, productObj.ProductType, "anggur")
	})

	t.Run("UpdateData | Invalid", func(t *testing.T) {
		productRepo.On("GetByID", mock.AnythingOfType("int")).Return(productDomain, nil).Once()
		productRepo.On("Update", mock.AnythingOfType("int"), productDomain).Return(domain.Product{}, errors.New("error")).Once()
		_, err := productService.UpdateData(productDomain.Id, productDomain)

		assert.Error(t, err)
	})
}

func TestDeleteData(t *testing.T) {
	productService = service.NewProductService(&productRepo)
	productDomain = domain.Product{
		Id:          1,
		UserId:      1,
		ProductType: "buah anggur",
	}

	t.Run("DeleteData | Valid", func(t *testing.T) {
		productRepo.On("Delete", mock.AnythingOfType("int")).Return(nil).Once()
		err := productService.DeleteData(productDomain.Id)

		assert.Nil(t, err)
	})

	t.Run("DeleteData | Invalid", func(t *testing.T) {
		productRepo.On("Delete", mock.AnythingOfType("int")).Return(errors.New("error")).Once()
		err := productService.DeleteData(productDomain.Id)

		assert.Error(t, err)
	})
}
