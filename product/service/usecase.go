package service

import (
	"errors"
	"fmt"

	errConv "github.com/Jiran03/gudhani/helper/error"
	"github.com/Jiran03/gudhani/product/domain"
)

type productService struct {
	repository domain.Repository
}

func (ps productService) InsertData(domain domain.Product) (productObj domain.Product, err error) {
	productObj, err = ps.repository.Create(domain)

	if err != nil {
		return domain, err
	}
	// errResp := errConv.Conversion(err)

	return productObj, nil
}

func (ps productService) GetAllData() (productObj []domain.Product, err error) {
	productObj, err = ps.repository.Get()

	if err != nil {
		return productObj, err
	}

	return productObj, nil
}

func (ps productService) GetDataByID(id int) (productObj domain.Product, err error) {
	productObj, err = ps.repository.GetByID(id)

	if err != nil {
		return productObj, errors.New(errConv.ErrDBNotFound)
	}

	return productObj, nil
}

func (ps productService) UpdateData(id int, domain domain.Product) (productObj domain.Product, err error) {
	product, errGetByID := ps.repository.GetByID(id)
	if errGetByID != nil {
		return domain, errGetByID
	}
	productId := product.Id
	productObj, err = ps.repository.Update(productId, domain)
	if err != nil {
		return domain, err
	}
	fmt.Println(productObj)
	return productObj, nil
}

func (ps productService) DeleteData(id int) (err error) {
	err = ps.repository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func NewProductService(repo domain.Repository) domain.Service {
	return productService{
		repository: repo,
	}
}
