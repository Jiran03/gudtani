//Repository hubungannya dengan data, data itu mau diapakan aja misal Save(), Update() dan sebagainya, Service untuk business logic misalkan createToken()
package domain

type Service interface {
	InsertData(domain Product) (productObj Product, err error)
	GetAllData() (productObj []Product, err error)
	GetDataByID(id int) (productObj Product, err error)
	UpdateData(id int, domain Product) (productObj Product, err error)
	DeleteData(id int) (err error)
}

type Repository interface {
	Create(domain Product) (productObj Product, err error)
	Update(id int, domain Product) (productObj Product, err error)
	Get() (productObj []Product, err error)
	GetByID(id int) (productObj Product, err error)
	Delete(id int) (err error)
}
