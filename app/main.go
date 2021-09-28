package main

import (
	"log"

	_driverFactory "laundro-api-ca/drivers"

	_addressRepo "laundro-api-ca/drivers/databases/addresses"

	_userService "laundro-api-ca/business/users"
	_userController "laundro-api-ca/controllers/users"
	_userRepo "laundro-api-ca/drivers/databases/users"

	_laundroService "laundro-api-ca/business/laundromats"
	_laundroController "laundro-api-ca/controllers/laundromats"
	_laundroRepo "laundro-api-ca/drivers/databases/laundromats"

	_dbDriver "laundro-api-ca/drivers/mysql"

	_middleware "laundro-api-ca/app/middleware"
	_routes "laundro-api-ca/app/routes"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func init() {
	viper.SetConfigFile(`app/config/config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func dbMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&_addressRepo.Addresses{},
		&_userRepo.Roles{},
		&_userRepo.Users{},
		&_laundroRepo.Laundromats{},
	)
	roles := []_userRepo.Roles{{ID: 1, Name: "Customer"},{ID: 2, Name: "Owner"}}
	db.Create(&roles)
	
}

func main() {
	configDB := _dbDriver.ConfigDB{
		DB_Username: viper.GetString(`database.user`),
		DB_Password: viper.GetString(`database.pass`),
		DB_Host:     viper.GetString(`database.host`),
		DB_Port:     viper.GetString(`database.port`),
		DB_Database: viper.GetString(`database.name`),
	}
	db := configDB.InitDB()
	dbMigrate(db)

	configJWT := _middleware.ConfigJWT{
		SecretJWT:       viper.GetString(`jwt.secret`),
		ExpiresDuration: int64(viper.GetInt(`jwt.expired`)),
	}

	e := echo.New()

	addrRepo := _driverFactory.NewAddressRepository(db)
	geoRepo := _driverFactory.NewGeolocationRepository()

	userRepo := _driverFactory.NewUserRepository(db)
	userService := _userService.NewUserService(userRepo, addrRepo, &configJWT)
	userCtrl := _userController.NewUserController(userService)
	
	laundroRepo := _driverFactory.NewLaundromatRepository(db)
	laundroService := _laundroService.NewLaundromatService(laundroRepo, addrRepo, geoRepo)
	laundroCtrl := _laundroController.NewLaundromatController(laundroService)

	routesInit := _routes.ControllerList{
		JWTMiddleware:        configJWT.Init(),
		UserController:       *userCtrl,
		LaundromatController: *laundroCtrl,
	}
	routesInit.RouteRegister(e)

	log.Fatal(e.Start(viper.GetString("server.address")))
}