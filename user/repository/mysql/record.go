package repository

import (
	"time"

	repoMySQLU "github.com/Jiran03/gudtani/product/repository/mysql"
	"github.com/Jiran03/gudtani/user/domain"
	repoMySQLW "github.com/Jiran03/gudtani/warehouse/repository/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID         int
	Name       string
	Email      string
	Password   string
	Address    string
	Gender     string
	Role       string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Products   []repoMySQLU.Product   `gorm:"foreignKey:UserID"`
	Warehouses []repoMySQLW.Warehouse `gorm:"foreignKey:UserID"`
}

func toDomain(rec User) domain.User {
	return domain.User{
		ID:        rec.ID,
		Name:      rec.Name,
		Email:     rec.Email,
		Password:  rec.Password,
		Address:   rec.Address,
		Gender:    rec.Gender,
		Role:      rec.Role,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
	}
}

func fromDomain(rec domain.User) User {
	return User{
		ID:        rec.ID,
		Name:      rec.Name,
		Email:     rec.Email,
		Password:  rec.Password,
		Address:   rec.Address,
		Gender:    rec.Gender,
		Role:      rec.Role,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
	}
}
