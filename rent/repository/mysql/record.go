package repository

import (
	"time"

	"github.com/Jiran03/gudhani/rent/domain"
	"gorm.io/gorm"
)

type Rent struct {
	gorm.Model
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

func toDomain(rec Rent) domain.Rent {
	return domain.Rent{
		ID:          rec.ID,
		ProductID:   rec.ProductID,
		WarehouseID: rec.WarehouseID,
		Weight:      rec.Weight,
		Period:      rec.Period,
		Status:      rec.Status,
		TotalPrice:  rec.TotalPrice,
		CreatedAt:   rec.CreatedAt,
		UpdatedAt:   rec.UpdatedAt,
	}
}

func fromDomain(rec domain.Rent) Rent {
	return Rent{
		ID:          rec.ID,
		ProductID:   rec.ProductID,
		WarehouseID: rec.WarehouseID,
		Weight:      rec.Weight,
		Period:      rec.Period,
		Status:      rec.Status,
		TotalPrice:  rec.TotalPrice,
	}
}

type RentalPrice struct {
	RentalPrice int
}

type Capacity struct {
	Capacity int
}
