package user

import (
	authMiddleware "github.com/Jiran03/gudtani/auth/middleware"
	userHandlerAPI "github.com/Jiran03/gudtani/user/handler/api"
	userRepositoryMySQL "github.com/Jiran03/gudtani/user/repository/mysql"
	userService "github.com/Jiran03/gudtani/user/service"
	"gorm.io/gorm"
)

func NewUserFactory(db *gorm.DB, configJWT authMiddleware.ConfigJWT) (userHandler userHandlerAPI.UserHandler) {
	userRepo := userRepositoryMySQL.NewUserRepository(db)
	userServ := userService.NewUserService(userRepo, configJWT)
	userHandler = userHandlerAPI.NewUserHandler(userServ)
	return
}
