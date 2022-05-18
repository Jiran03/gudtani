package warehouse

import (
	warehouseHandlerAPI "github.com/Jiran03/gudhani/warehouse/handler/api"
	warehouseRepositoryMySQL "github.com/Jiran03/gudhani/warehouse/repository/mysql"
	warehouseService "github.com/Jiran03/gudhani/warehouse/service"
	"gorm.io/gorm"
)

func NewWarehouseFactory(db *gorm.DB) (warehouseHandler warehouseHandlerAPI.WarehouseHandler) {
	warehouseRepo := warehouseRepositoryMySQL.NewWarehouseRepository(db)
	warehouseServ := warehouseService.NewWarehouseService(warehouseRepo)
	warehouseHandler = warehouseHandlerAPI.NewWarehouseHandler(warehouseServ)
	return
}
