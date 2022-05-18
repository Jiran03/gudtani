package domain

import "time"

type Warehouse struct {
	Id            int
	UserId        int
	WarehouseName string
	Capacity      int
	RentalPrice   int
	Address       string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type WarehouseDescription struct {
	WarehouseId int
	Description string
}
