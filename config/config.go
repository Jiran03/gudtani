package config

import (
	"fmt"

	repoProduct "github.com/Jiran03/gudhani/product/repository/mysql"
	repoRent "github.com/Jiran03/gudhani/rent/repository/mysql"
	repoUser "github.com/Jiran03/gudhani/user/repository/mysql"
	repoWarehouse "github.com/Jiran03/gudhani/warehouse/repository/mysql"
	"github.com/spf13/viper"
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
		DBNAME: viper.GetString(`DBNAME`),
		DBUSER: viper.GetString(`DBUSER`),
		DBPASS: viper.GetString(`DBPASS`),
		DBHOST: viper.GetString(`DBHOST`),
		DBPORT: viper.GetString(`DBPORT`),
		// JWTSecret: os.Getenv("JWTSECRET"),
		// DBNAME: "gudhani",
		// DBUSER: "root",
		// DBPASS: "",
		// DBHOST: "localhost",
		// DBPORT: "3306",
	}
	fmt.Printf("%+v", Conf)
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

// //with .env file
// func (config *Config) InitDB() *gorm.DB {
// 	DB, _ := gorm.Open(
// 		mysql.Open(
// 			fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
// 				config.DBUSER,
// 				config.DBPASS,
// 				config.DBHOST,
// 				config.DBPORT,
// 				config.DBNAME,
// 			),
// 		),
// 	)
// 	return DB
// }

func DBMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&repoUser.User{},
		&repoProduct.Product{},
		&repoWarehouse.Warehouse{},
		&repoRent.Rent{},
	)
}
