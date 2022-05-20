package repository

import (
	"time"

	"github.com/Jiran03/gudtani/product/domain"
	repoMySQLR "github.com/Jiran03/gudtani/rent/repository/mysql"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ID          int
	UserID      int
	ProductType string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Rents       []repoMySQLR.Rent
}

func toDomain(rec Product) domain.Product {
	return domain.Product{
		Id:          rec.ID,
		UserId:      rec.UserID,
		ProductType: rec.ProductType,
		CreatedAt:   rec.CreatedAt,
		UpdatedAt:   rec.UpdatedAt,
	}
}

func fromDomain(rec domain.Product) Product {
	return Product{
		ID:          rec.Id,
		UserID:      rec.UserId,
		ProductType: rec.ProductType,
		CreatedAt:   rec.CreatedAt,
		UpdatedAt:   rec.UpdatedAt,
	}
}
