package domain

import "time"

type Rent struct {
	ID          int
	ProductID   int
	WarehouseID int
	Weight      int
	Period      int
	Status      string
	TotalPrice  int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
