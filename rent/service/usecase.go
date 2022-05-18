package service

import (
	"errors"

	errConv "github.com/Jiran03/gudhani/helper/error"
	"github.com/Jiran03/gudhani/rent/domain"
)

type rentService struct {
	repository domain.Repository
}

func (rs rentService) InsertData(domain domain.Rent) (rentObj domain.Rent, err error) {
	domain.Status = "sedang dalam penyewaan"

	if domain.Period == 0 {
		domain.Status = "masa sewa telah selesai"
	}

	rentalPrice, errGetRentalPrice := rs.repository.GetRentalPrice(domain.WarehouseID)

	if errGetRentalPrice != nil {
		return rentObj, errGetRentalPrice
	}

	totalPrice := rentalPrice * domain.Weight * domain.Period
	domain.TotalPrice = totalPrice
	rentObj, err = rs.repository.Create(domain)

	if err != nil {
		return domain, err
	}

	// errResp := errConv.Conversion(err)
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
		return rentObj, errors.New(errConv.ErrDBNotFound)
	}
	return rentObj, nil
}

func (rs rentService) UpdateData(id int, domain domain.Rent) (rentObj domain.Rent, err error) {
	rent, errGetByID := rs.repository.GetByID(id)
	if errGetByID != nil {
		return domain, errGetByID
	}

	domain.Status = "sedang dalam penyewaan"
	domain.TotalPrice = rent.TotalPrice
	domain.CreatedAt = rent.CreatedAt

	if domain.Period == 0 {
		domain.Status = "masa sewa telah selesai"
	}

	rentId := rent.ID
	rentObj, err = rs.repository.Update(rentId, domain)
	if err != nil {
		return domain, err
	}

	return rentObj, nil
}

func (rs rentService) DeleteData(id int) (err error) {
	err = rs.repository.Delete(id)
	if err != nil {
		return errConv.Conversion(err)
	}
	return nil
}

func NewRentService(repo domain.Repository) domain.Service {
	return rentService{
		repository: repo,
	}
}
