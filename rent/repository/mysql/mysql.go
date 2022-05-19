package repository

import (
	"github.com/Jiran03/gudtani/rent/domain"
	"gorm.io/gorm"
)

type rentRepo struct {
	DB *gorm.DB
}

func (rr rentRepo) Create(domain domain.Rent) (rentObj domain.Rent, err error) {
	record := fromDomain(domain)
	err = rr.DB.Create(&record).Error

	if err != nil {
		return domain, err
	}

	return toDomain(record), nil
}

func (rr rentRepo) Update(id int, domain domain.Rent) (rentObj domain.Rent, err error) {
	var newRecord Rent
	record := fromDomain(domain)
	err = rr.DB.Model(&newRecord).Where("id = ?", id).Updates(map[string]interface{}{
		"id":           id,
		"product_id":   record.ProductID,
		"warehouse_id": record.WarehouseID,
		"weight":       record.Weight,
		"period":       record.Period,
		"status":       record.Status,
		"total_price":  record.TotalPrice,
	}).Error

	if err != nil {
		return domain, err
	}

	rentObj = toDomain(newRecord)
	return rentObj, nil
}

func (rr rentRepo) Get() (rentObj []domain.Rent, err error) {
	var newRecord []Rent
	err = rr.DB.Find(&newRecord).Error

	if err != nil {
		return rentObj, err
	}

	for _, value := range newRecord {
		rentObj = append(rentObj, toDomain(value))
	}

	return rentObj, nil
}

func (rr rentRepo) GetByID(id int) (userObj domain.Rent, err error) {
	var newRecord Rent
	err = rr.DB.Where("id = ?", id).First(&newRecord).Error

	if err != nil {
		return userObj, err
	}

	return toDomain(newRecord), nil
}

func (rr rentRepo) Delete(id int) (err error) {
	var record Rent
	return rr.DB.Delete(&record, id).Error
}

func NewRentRepository(db *gorm.DB) domain.Repository {
	return rentRepo{
		DB: db,
	}
}
