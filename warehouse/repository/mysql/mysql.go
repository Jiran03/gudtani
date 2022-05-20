package repository

import (
	"github.com/Jiran03/gudtani/warehouse/domain"
	"gorm.io/gorm"
)

type warehouseRepo struct {
	DB *gorm.DB
}

func (wr warehouseRepo) Create(domain domain.Warehouse) (warehouseObj domain.Warehouse, err error) {
	record := fromDomain(domain)
	err = wr.DB.Create(&record).Error

	if err != nil {
		return domain, err
	}

	return toDomain(record), nil
}

func (wr warehouseRepo) Update(id int, domain domain.Warehouse) (warehouseObj domain.Warehouse, err error) {
	var newRecord Warehouse
	record := fromDomain(domain)
	err = wr.DB.Model(&newRecord).Where("id = ?", id).Updates(map[string]interface{}{
		"id":             id,
		"user_id":        record.UserID,
		"warehouse_name": record.WarehouseName,
		"capacity":       record.Capacity,
		"rental_price":   record.RentalPrice,
		"address":        record.Address,
	}).Error

	if err != nil {
		return domain, err
	}

	warehouseObj = toDomain(newRecord)
	return warehouseObj, nil
}

func (wr warehouseRepo) UpdateCapacity(id, newCapacity int) (err error) {
	err = wr.DB.Table("warehouses").Where("id = ?", id).Updates(map[string]interface{}{
		"capacity": newCapacity,
	}).Error

	if err != nil {
		return err
	}

	return nil
}

func (wr warehouseRepo) Get() (warehouseObj []domain.Warehouse, err error) {
	var newRecord []Warehouse
	err = wr.DB.Find(&newRecord).Error

	if err != nil {
		return warehouseObj, err
	}

	for _, value := range newRecord {
		warehouseObj = append(warehouseObj, toDomain(value))
	}

	return warehouseObj, nil
}

func (wr warehouseRepo) GetByID(id int) (warehouseObj domain.Warehouse, err error) {
	var newRecord Warehouse
	err = wr.DB.Where("id = ?", id).First(&newRecord).Error

	if err != nil {
		return warehouseObj, err
	}

	return toDomain(newRecord), nil
}

func (wr warehouseRepo) GetByAddress(address string) (warehouseObj []domain.Warehouse, err error) {
	var newRecord []Warehouse
	err = wr.DB.Where("address LIKE ? AND capacity > ?", "%"+address+"%", 0).Find(&newRecord).Error

	if err != nil {
		return warehouseObj, err
	}

	for _, value := range newRecord {
		warehouseObj = append(warehouseObj, toDomain(value))
	}

	return warehouseObj, nil
}

func (wr warehouseRepo) Delete(id int) (err error) {
	var record Warehouse
	return wr.DB.Delete(&record, id).Error
}

func NewWarehouseRepository(db *gorm.DB) domain.Repository {
	return warehouseRepo{
		DB: db,
	}
}
