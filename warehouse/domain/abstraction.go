//Repository hubungannya dengan data, data itu mau diapakan aja misal Save(), Update() dan sebagainya, Service untuk business logic misalkan createToken()

package domain

type Service interface {
	InsertData(domain Warehouse) (warehouseObj Warehouse, err error)
	GetAllData() (warehouseObj []Warehouse, err error)
	GetDataByID(id int) (warehouseObj Warehouse, err error)
	GetDataByAddress(address string) (warehouseObj []Warehouse, err error)
	UpdateData(id int, domain Warehouse) (warehouseObj Warehouse, err error)
	UpdateDataCapacity(id, newCapacity int) (err error)
	DeleteData(id int) (err error)
}

type Repository interface {
	Create(domain Warehouse) (warehouseObj Warehouse, err error)
	Update(id int, domain Warehouse) (warehouseObj Warehouse, err error)
	Get() (warehouseObj []Warehouse, err error)
	GetByAddress(address string) (warehouseObj []Warehouse, err error)
	GetByID(id int) (warehouseObj Warehouse, err error)
	UpdateCapacity(id, newCapacity int) (err error)
	Delete(id int) (err error)
}
