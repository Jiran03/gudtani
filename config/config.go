package config

import (
	"fmt"
	"os"

	repoProduct "github.com/Jiran03/gudtani/product/repository/mysql"
	repoRent "github.com/Jiran03/gudtani/rent/repository/mysql"
	repoUser "github.com/Jiran03/gudtani/user/repository/mysql"
	repoWarehouse "github.com/Jiran03/gudtani/warehouse/repository/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	DBNAME string
	DBUSER string
	DBPASS string
	DBHOST string
	DBPORT string
}

var Conf Config

func Init() {
	Conf = Config{
		DBNAME: os.Getenv("DBNAME"),
		DBUSER: os.Getenv("DBUSER"),
		DBPASS: os.Getenv("DBPASS"),
		DBHOST: os.Getenv("DBHOST"),
		DBPORT: os.Getenv("DBPORT"),
	}
}

func DBInit() (DB *gorm.DB) {
	DB, _ = gorm.Open(
		mysql.Open(
			fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
				Conf.DBUSER,
				Conf.DBPASS,
				Conf.DBHOST,
				Conf.DBPORT,
				Conf.DBNAME,
			),
		),
	)
	return
}

func DBMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&repoUser.User{},
		&repoProduct.Product{},
		&repoWarehouse.Warehouse{},
		&repoRent.Rent{},
	)
}
