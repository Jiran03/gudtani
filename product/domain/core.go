package domain

import "time"

type Product struct {
	Id          int
	UserId      int
	ProductType string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
