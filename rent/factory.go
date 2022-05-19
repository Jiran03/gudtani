package rent

import (
	rentHandlerAPI "github.com/Jiran03/gudtani/rent/handler/api"
	rentRepositoryMySQL "github.com/Jiran03/gudtani/rent/repository/mysql"
	rentService "github.com/Jiran03/gudtani/rent/service"
	warehouseRepositoryMySQL "github.com/Jiran03/gudtani/warehouse/repository/mysql"
	warehouseService "github.com/Jiran03/gudtani/warehouse/service"
	"gorm.io/gorm"
)

func NewRentFactory(db *gorm.DB) (rentHandler rentHandlerAPI.RentHandler) {
	warehouseRepo := warehouseRepositoryMySQL.NewWarehouseRepository(db)
	warehouseServ := warehouseService.NewWarehouseService(warehouseRepo)
	rentRepo := rentRepositoryMySQL.NewRentRepository(db)
	rentServ := rentService.NewRentService(rentRepo, warehouseServ)
	rentHandler = rentHandlerAPI.NewRentHandler(rentServ)
	return
}
