package service

import (
	"github.com/Jiran03/gudtani/warehouse/domain"
)

type warehouseService struct {
	repository domain.Repository
}

func (ws warehouseService) InsertData(domain domain.Warehouse) (warehouseObj domain.Warehouse, err error) {
	warehouseObj, err = ws.repository.Create(domain)

	if err != nil {
		return domain, err
	}

	return warehouseObj, nil
}
func (ws warehouseService) GetAllData() (warehouseObj []domain.Warehouse, err error) {
	warehouseObj, err = ws.repository.Get()

	if err != nil {
		return warehouseObj, err
	}

	return warehouseObj, nil
}

func (ws warehouseService) GetDataByID(id int) (warehouseObj domain.Warehouse, err error) {
	warehouseObj, err = ws.repository.GetByID(id)

	if err != nil {
		return warehouseObj, err
	}

	return warehouseObj, nil
}

func (ws warehouseService) GetDataByAddress(address string) (warehouseObj []domain.Warehouse, err error) {
	warehouseObj, err = ws.repository.GetByAddress(address)

	if err != nil {
		return warehouseObj, err
	}

	return warehouseObj, nil
}

func (ws warehouseService) UpdateData(id int, domain domain.Warehouse) (warehouseObj domain.Warehouse, err error) {
	warehouse, errGetByID := ws.repository.GetByID(id)

	if errGetByID != nil {
		return domain, errGetByID
	}

	warehouseId := warehouse.Id
	warehouseObj, err = ws.repository.Update(warehouseId, domain)

	if err != nil {
		return domain, err
	}

	return warehouseObj, nil
}

func (ws warehouseService) UpdateDataCapacity(id, newWarehouseCapacity int) (err error) {
	err = ws.repository.UpdateCapacity(id, newWarehouseCapacity)

	if err != nil {
		return err
	}

	return nil
}

func (ws warehouseService) DeleteData(id int) (err error) {
	err = ws.repository.Delete(id)

	if err != nil {
		return err
	}

	return nil
}

func NewWarehouseService(repo domain.Repository) domain.Service {
	return warehouseService{
		repository: repo,
	}
}
