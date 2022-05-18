package rent

import (
	rentHandlerAPI "github.com/Jiran03/gudhani/rent/handler/api"
	rentRepositoryMySQL "github.com/Jiran03/gudhani/rent/repository/mysql"
	rentService "github.com/Jiran03/gudhani/rent/service"
	"gorm.io/gorm"
)

func NewRentFactory(db *gorm.DB) (rentHandler rentHandlerAPI.RentHandler) {
	rentRepo := rentRepositoryMySQL.NewRentRepository(db)
	rentServ := rentService.NewRentService(rentRepo)
	rentHandler = rentHandlerAPI.NewRentHandler(rentServ)
	return
}
