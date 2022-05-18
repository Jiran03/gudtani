package repository

import (
	"fmt"

	"github.com/Jiran03/gudhani/product/domain"
	"gorm.io/gorm"
)

type productRepo struct {
	DB *gorm.DB
}

func (pr productRepo) Create(domain domain.Product) (productObj domain.Product, err error) {
	record := fromDomain(domain)
	err = pr.DB.Create(&record).Error
	if err != nil {
		return domain, err
	}
	return toDomain(record), nil
}

func (pr productRepo) Update(id int, domain domain.Product) (productObj domain.Product, err error) {
	var newRecord Product
	fmt.Println("id repo", id)
	record := fromDomain(domain)
	err = pr.DB.Model(&newRecord).Where("id = ?", id).Updates(map[string]interface{}{
		"id":           id,
		"user_id":      record.UserID,
		"product_type": record.ProductType,
	}).Error

	if err != nil {
		return domain, err
	}

	productObj = toDomain(newRecord)
	return productObj, nil
}

func (pr productRepo) Get() (productObj []domain.Product, err error) {
	var newRecord []Product
	err = pr.DB.Find(&newRecord).Error
	if err != nil {
		return productObj, err
	}

	for _, value := range newRecord {
		productObj = append(productObj, toDomain(value))
	}

	return productObj, nil
}

func (pr productRepo) GetByID(id int) (userObj domain.Product, err error) {
	var newRecord Product
	err = pr.DB.Where("id = ?", id).First(&newRecord).Error
	if err != nil {
		return userObj, err
	}
	return toDomain(newRecord), nil
}

func (pr productRepo) Delete(id int) (err error) {
	var record Product
	return pr.DB.Delete(&record, id).Error
}

func NewProductRepository(db *gorm.DB) domain.Repository {
	return productRepo{
		DB: db,
	}
}
