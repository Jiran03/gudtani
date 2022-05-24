package main

import (
	"log"

	"github.com/Jiran03/gudtani/auth"
	authMiddleware "github.com/Jiran03/gudtani/auth/middleware"
	"github.com/Jiran03/gudtani/auth/routes"
	"github.com/Jiran03/gudtani/config"
	"github.com/Jiran03/gudtani/product"
	"github.com/Jiran03/gudtani/rent"
	"github.com/Jiran03/gudtani/user"
	"github.com/Jiran03/gudtani/warehouse"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func init() {
	err := godotenv.Load("config.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	config.Init()
	db := config.DBInit()
	config.DBMigrate(db)

	configJWT := authMiddleware.ConfigJWT{
		SecretJWT:       auth.SECRET_KEY,
		ExpiresDuration: auth.EXPIRED,
	}

	user := user.NewUserFactory(db, configJWT)
	product := product.NewProductFactory(db)
	warehouse := warehouse.NewWarehouseFactory(db)
	rent := rent.NewRentFactory(db)

	e := echo.New()
	authMiddleware.LogMiddlewares(e)
	cJWT := configJWT.Init()

	e.POST("/register", user.Register)
	e.POST("/login", user.Login)
	e.POST("/search", warehouse.GetDataByAddress)

	userGroup := e.Group("/user")
	userGroup.Use(middleware.JWTWithConfig(cJWT))
	userGroup.GET("", user.GetAllData)
	userGroup.GET("/:id", user.GetByID)
	userGroup.PUT("/:id", user.Update)
	userGroup.DELETE("/:id", user.Delete)

	productGroup := e.Group("/product")
	productGroup.Use(middleware.JWTWithConfig(cJWT), routes.RoleValidation("petani", user))
	productGroup.POST("", product.InsertData)
	productGroup.GET("", product.GetAllData)
	productGroup.GET("/:id", product.GetDataByID)
	productGroup.PUT("/:id", product.UpdateData)
	productGroup.DELETE("/:id", product.DeleteData)

	warehouseGroup := e.Group("/warehouse")
	warehouseGroup.Use(middleware.JWTWithConfig(cJWT), routes.RoleValidation("pemilik gudang", user))
	warehouseGroup.POST("", warehouse.InsertData)
	warehouseGroup.GET("", warehouse.GetAllData)
	warehouseGroup.GET("/:id", warehouse.GetDataByID)
	warehouseGroup.PUT("/:id", warehouse.UpdateData)
	warehouseGroup.DELETE("/:id", warehouse.DeleteData)

	rentGroup := e.Group("/rent")
	rentGroup.Use(middleware.JWTWithConfig(cJWT))
	rentGroup.POST("", rent.InsertData)
	rentGroup.GET("", rent.GetAllData)
	rentGroup.GET("/:id", rent.GetDataByID)
	rentGroup.PUT("/:id", rent.UpdateData)
	rentGroup.DELETE("/:id", rent.DeleteData)

	e.Start(":9500")
}
