package service

import (
	"errors"

	"github.com/Jiran03/gudtani/rent/domain"
	warehouseDomain "github.com/Jiran03/gudtani/warehouse/domain"
)

type rentService struct {
	repository    domain.Repository
	warehouseServ warehouseDomain.Service
}

func (rs rentService) InsertData(domain domain.Rent) (rentObj domain.Rent, err error) {
	domain.Status = "sedang dalam penyewaan"

	if domain.Period == 0 {
		domain.Status = "masa sewa telah selesai"
	}

	warehouseID := domain.WarehouseID
	warehouseObj, err := rs.warehouseServ.GetDataByID(warehouseID)

	if err != nil {
		return rentObj, err
	}

	if warehouseObj.Capacity == 0 || warehouseObj.Capacity < domain.Weight {
		return rentObj, errors.New("warehouse capacity is not enough")
	}

	newWarehouseCapacity := warehouseObj.Capacity - domain.Weight
	errUpdateCapacity := rs.warehouseServ.UpdateDataCapacity(warehouseID, newWarehouseCapacity)

	if errUpdateCapacity != nil {
		return rentObj, errUpdateCapacity
	}

	rentalPrice := warehouseObj.RentalPrice
	totalPrice := rentalPrice * domain.Weight * domain.Period
	domain.TotalPrice = totalPrice
	rentObj, err = rs.repository.Create(domain)

	if err != nil {
		return domain, err
	}

	return rentObj, nil
}
func (rs rentService) GetAllData() (rentObj []domain.Rent, err error) {
	rentObj, err = rs.repository.Get()

	if err != nil {
		return rentObj, err
	}

	return rentObj, nil
}

func (rs rentService) GetDataByID(id int) (rentObj domain.Rent, err error) {
	rentObj, err = rs.repository.GetByID(id)

	if err != nil {
		return rentObj, err
	}

	return rentObj, nil
}

func (rs rentService) UpdateData(id int, domain domain.Rent) (rentObj domain.Rent, err error) {
	rent, errGetByID := rs.repository.GetByID(id)

	if errGetByID != nil {
		return domain, errGetByID
	}

	warehouseID := domain.WarehouseID
	warehouseObj, err := rs.warehouseServ.GetDataByID(warehouseID)

	if err != nil {
		return rentObj, err
	}

	newWarehouseCapacity := warehouseObj.Capacity
	domain.Status = "sedang dalam penyewaan"

	if domain.Period == 0 {
		newWarehouseCapacity += rent.Weight
		domain.Status = "masa sewa telah selesai"
	}

	errUpdateCapacity := rs.warehouseServ.UpdateDataCapacity(warehouseID, newWarehouseCapacity)

	if errUpdateCapacity != nil {
		return rentObj, errUpdateCapacity
	}

	domain.TotalPrice = rent.TotalPrice
	rentObj, err = rs.repository.Update(id, domain)

	if err != nil {
		return domain, err
	}

	return rentObj, nil
}

func (rs rentService) DeleteData(id int) (err error) {
	err = rs.repository.Delete(id)

	if err != nil {
		return err
	}

	return nil
}

func NewRentService(repo domain.Repository, ws warehouseDomain.Service) domain.Service {
	return rentService{
		repository:    repo,
		warehouseServ: ws,
	}
}
