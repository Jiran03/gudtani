//Repository hubungannya dengan data, data itu mau diapakan aja misal Save(), Update() dan sebagainya, Service untuk business logic misalkan createToken()

package domain

type Service interface {
	// CreateToken(id int, role string) (token string, err error)
	CreateToken(email, password string) (token string, err error)
	InsertData(domain User) (userObj User, err error)
	GetAllData() (userObj []User, err error)
	GetByID(id int) (userObj User, err error)
	GetByEmailPassword(email, password string) (id int, role string, err error)
	UpdateData(id int, domain User) (userObj User, err error)
	DeleteData(id int) (err error)
}

type Repository interface {
	Create(domain User) (userObj User, err error)
	Update(id int, domain User) (userObj User, err error)
	Delete(id int) (err error)
	Get() (userObj []User, err error)
	GetByEmailPassword(email, password string) (domain User, err error)
	GetByID(id int) (userObj User, err error)
}
