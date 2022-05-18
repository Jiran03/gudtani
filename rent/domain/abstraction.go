//Repository hubungannya dengan data, data itu mau diapakan aja misal Save(), Update() dan sebagainya, Service untuk business logic misalkan createToken()

package domain

type Service interface {
	InsertData(domain Rent) (rentObj Rent, err error)
	GetAllData() (rentObj []Rent, err error)
	GetDataByID(id int) (rentObj Rent, err error)
	UpdateData(id int, domain Rent) (rentObj Rent, err error)
	DeleteData(id int) (err error)
}

type Repository interface {
	Create(domain Rent) (rentObj Rent, err error)
	Update(id int, domain Rent) (rentObj Rent, err error)
	GetRentalPrice(id int) (rentalPrice int, err error)
	Get() (rentObj []Rent, err error)
	GetByID(id int) (rentObj Rent, err error)
	Delete(id int) (err error)
}
