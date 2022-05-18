package repository

import (
	"time"

	repoMySQLR "github.com/Jiran03/gudhani/rent/repository/mysql"
	"github.com/Jiran03/gudhani/warehouse/domain"
	"gorm.io/gorm"
)

type Warehouse struct {
	gorm.Model
	ID            int
	UserID        int
	WarehouseName string
	Capacity      int
	RentalPrice   int
	Address       string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Rents         []repoMySQLR.Rent
}

func toDomain(rec Warehouse) domain.Warehouse {
	return domain.Warehouse{
		Id:            rec.ID,
		UserId:        rec.UserID,
		WarehouseName: rec.WarehouseName,
		Capacity:      rec.Capacity,
		RentalPrice:   rec.RentalPrice,
		Address:       rec.Address,
		CreatedAt:     rec.CreatedAt,
		UpdatedAt:     rec.UpdatedAt,
	}
}

func fromDomain(rec domain.Warehouse) Warehouse {
	return Warehouse{
		ID:            rec.Id,
		UserID:        rec.UserId,
		WarehouseName: rec.WarehouseName,
		Capacity:      rec.Capacity,
		RentalPrice:   rec.RentalPrice,
		Address:       rec.Address,
		CreatedAt:     rec.CreatedAt,
		UpdatedAt:     rec.UpdatedAt,
	}
}
