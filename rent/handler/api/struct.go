package rentAPI

import (
	"time"

	"github.com/Jiran03/gudhani/rent/domain"
)

type RequestJSON struct {
	ProductID   int `json:"product_id"`
	WarehouseID int `json:"warehouse_id"`
	Weight      int `json:"weight"`
	Period      int `json:"period"`
}

func toDomain(req RequestJSON) domain.Rent {
	return domain.Rent{
		ProductID:   req.ProductID,
		WarehouseID: req.WarehouseID,
		Weight:      req.Weight,
		Period:      req.Period,
	}
}

type ResponseJSON struct {
	Id          int       `json:"id"`
	ProductID   int       `json:"product_id"`
	WarehouseID int       `json:"warehouse_id"`
	Weight      int       `json:"weight"`
	Period      int       `json:"period"`
	Status      string    `json:"status"`
	TotalPrice  int       `json:"total_price"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func fromDomain(domain domain.Rent) ResponseJSON {
	return ResponseJSON{
		Id:          domain.ID,
		ProductID:   domain.ProductID,
		WarehouseID: domain.WarehouseID,
		Weight:      domain.Weight,
		Period:      domain.Period,
		Status:      domain.Status,
		TotalPrice:  domain.TotalPrice,
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
	}
}
