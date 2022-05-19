package product

import (
	productHandlerAPI "github.com/Jiran03/gudtani/product/handler/api"
	productRepositoryMySQL "github.com/Jiran03/gudtani/product/repository/mysql"
	productService "github.com/Jiran03/gudtani/product/service"
	"gorm.io/gorm"
)

func NewProductFactory(db *gorm.DB) (productHandler productHandlerAPI.ProductHandler) {
	productRepo := productRepositoryMySQL.NewProductRepository(db)
	productServ := productService.NewProductService(productRepo)
	productHandler = productHandlerAPI.NewProductHandler(productServ)
	return
}
