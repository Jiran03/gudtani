package productAPI

import (
	"time"

	"github.com/Jiran03/gudtani/product/domain"
)

type RequestJSON struct {
	UserId      int    `json:"user_id" validate:"required"`
	ProductType string `json:"product_type" validate:"required"`
}

func toDomain(req RequestJSON) domain.Product {
	return domain.Product{
		UserId:      req.UserId,
		ProductType: req.ProductType,
	}
}

type ResponseJSON struct {
	Id          int       `json:"id"`
	UserId      int       `json:"user_id"`
	ProductType string    `json:"product_type"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func fromDomain(domain domain.Product) ResponseJSON {
	return ResponseJSON{
		Id:          domain.Id,
		UserId:      domain.UserId,
		ProductType: domain.ProductType,
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
	}
}
