package userAPI

import (
	"time"

	"github.com/Jiran03/gudtani/user/domain"
)

type RequestJSON struct {
	Name     string `json:"name"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Address  string `json:"address" validate:"required"`
	Gender   string `json:"gender" validate:"required"`
	Role     string `json:"role" validate:"required"`
}

type RequestLoginJSON struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func toDomain(req RequestJSON) domain.User {
	return domain.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
		Address:  req.Address,
		Gender:   req.Gender,
		Role:     req.Role,
	}
}

type ResponseJSON struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Address   string    `json:"address"`
	Gender    string    `json:"gender"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func fromDomain(domain domain.User) ResponseJSON {
	return ResponseJSON{
		Id:        domain.ID,
		Name:      domain.Name,
		Email:     domain.Email,
		Address:   domain.Address,
		Gender:    domain.Gender,
		Role:      domain.Role,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
