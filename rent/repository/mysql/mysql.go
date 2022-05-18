package repository

import (
	"errors"
	"fmt"

	"github.com/Jiran03/gudhani/rent/domain"
	"gorm.io/gorm"
)

type rentRepo struct {
	DB *gorm.DB
}

func (rr rentRepo) Create(domain domain.Rent) (rentObj domain.Rent, err error) {
	var warehouseCapacity Capacity
	record := fromDomain(domain)

	fmt.Println(record.WarehouseID)

	//update data warehouse capacity
	errJoinTable := rr.DB.Table("warehouses").Select("warehouses.capacity").Where("id = ?", record.WarehouseID).First(&warehouseCapacity).Error

	if errJoinTable != nil {
		return rentObj, errJoinTable
	}

	if warehouseCapacity.Capacity == 0 || warehouseCapacity.Capacity < record.Weight {
		return rentObj, errors.New("warehouse capacity is not enough")
	}

	newWarehouseCapacity := warehouseCapacity.Capacity - record.Weight
	errUpdateWarehouseCapacity := rr.DB.Table("warehouses").Where("id = ?", record.WarehouseID).Updates(map[string]interface{}{
		"capacity": newWarehouseCapacity,
	}).Error

	if errUpdateWarehouseCapacity != nil {
		return rentObj, errUpdateWarehouseCapacity
	}

	//create new rent
	err = rr.DB.Create(&record).Error

	if err != nil {
		return domain, err
	}

	return toDomain(record), nil
}

func (rr rentRepo) Update(id int, domain domain.Rent) (rentObj domain.Rent, err error) {
	var newRecord Rent
	var warehouseCapacity Capacity
	var newWarehouseCapacity int
	record := fromDomain(domain)

	//update data warehouse capacity
	errJoinTable := rr.DB.Table("warehouses").Select("warehouses.capacity").Where("id = ?", record.WarehouseID).First(&warehouseCapacity).Error

	if errJoinTable != nil {
		return rentObj, errJoinTable
	}

	if record.Period == 0 {
		newWarehouseCapacity = warehouseCapacity.Capacity + record.Weight
	}

	errUpdateWarehouseCapacity := rr.DB.Table("warehouses").Where("id = ?", record.WarehouseID).Updates(map[string]interface{}{
		"capacity": newWarehouseCapacity,
	}).Error

	if errUpdateWarehouseCapacity != nil {
		return rentObj, errUpdateWarehouseCapacity
	}

	// update rent
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

func (rr rentRepo) GetRentalPrice(id int) (rentalPrice int, err error) {
	var result RentalPrice
	err = rr.DB.Table("warehouses").Select("warehouses.rental_price").Where("id = ?", id).First(&result).Error

	if err != nil {
		return 0, err
	}

	return result.RentalPrice, nil
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
