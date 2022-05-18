package warehouseAPI

import (
	"time"

	"github.com/Jiran03/gudhani/warehouse/domain"
)

type RequestJSON struct {
	UserId        int    `json:"user_id" validate:"required"`
	WarehouseName string `json:"warehouse_name" validate:"required"`
	Capacity      int    `json:"capacity" validate:"required"`
	RentalPrice   int    `json:"rental_price" validate:"required"`
	Address       string `json:"address" validate:"required"`
}

func toDomain(req RequestJSON) domain.Warehouse {
	return domain.Warehouse{
		UserId:        req.UserId,
		WarehouseName: req.WarehouseName,
		Capacity:      req.Capacity,
		RentalPrice:   req.RentalPrice,
		Address:       req.Address,
	}
}

type ResponseJSON struct {
	Id            int       `json:"id"`
	UserId        int       `json:"user_id"`
	WarehouseName string    `json:"warehouse_name"`
	Capacity      int       `json:"capacity"`
	RentalPrice   int       `json:"rental_price"`
	Address       string    `json:"address"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func fromDomain(domain domain.Warehouse) ResponseJSON {
	return ResponseJSON{
		Id:            domain.Id,
		UserId:        domain.UserId,
		WarehouseName: domain.WarehouseName,
		Capacity:      domain.Capacity,
		RentalPrice:   domain.RentalPrice,
		Address:       domain.Address,
		CreatedAt:     domain.CreatedAt,
		UpdatedAt:     domain.UpdatedAt,
	}
}
