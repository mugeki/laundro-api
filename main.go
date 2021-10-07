package main

import (
	"log"
	"os"
	"strconv"

	_driverFactory "laundro-api-ca/drivers"

	_addressRepo "laundro-api-ca/drivers/databases/addresses"

	_userService "laundro-api-ca/business/users"
	_userController "laundro-api-ca/controllers/users"
	_userRepo "laundro-api-ca/drivers/databases/users"

	_laundroService "laundro-api-ca/business/laundromats"
	_laundroController "laundro-api-ca/controllers/laundromats"
	_laundroRepo "laundro-api-ca/drivers/databases/laundromats"

	_productService "laundro-api-ca/business/products"
	_productController "laundro-api-ca/controllers/products"
	_productRepo "laundro-api-ca/drivers/databases/products"

	_orderService "laundro-api-ca/business/orders"
	_orderController "laundro-api-ca/controllers/orders"
	_orderRepo "laundro-api-ca/drivers/databases/orders"

	_dbDriver "laundro-api-ca/drivers/mysql"

	"laundro-api-ca/app/middleware"
	_middleware "laundro-api-ca/app/middleware"
	_routes "laundro-api-ca/app/routes"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func dbMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&_addressRepo.Addresses{},
		&_userRepo.Roles{},
		&_userRepo.Users{},
		&_productRepo.Products{},
		&_laundroRepo.Laundromats{},
		&_orderRepo.Orders{},
	)
	roles := []_userRepo.Roles{{ID: 1, Name: "Customer"},{ID: 2, Name: "Owner"}}
	categories := []_productRepo.Category{{ID: 1, Name: "Kiloan"},{ID: 2, Name: "Dry Clean"},{ID: 3, Name: "Cuci Sepatu"}}
	payments := []_orderRepo.Payment{{ID: 1, Gateway: "Gopay"},{ID: 2, Gateway: "OVO"}}
	db.Create(&roles)
	db.Create(&categories)
	db.Create(&payments)
}

func main() {
	_ = godotenv.Load()

	configDB := _dbDriver.ConfigDB{
		DB_Username: os.Getenv("DB_USER"),
		DB_Password: os.Getenv("DB_PASSWORD"),
		DB_Host:     os.Getenv("DB_HOST"),
		DB_Port:     os.Getenv("DB_PORT"),
		DB_Database: os.Getenv("DB_NAME"),
	}
	db := configDB.InitDB()
	dbMigrate(db)

	log.Println(configDB)

	EXPIRE, _ := strconv.Atoi(os.Getenv("JWT_EXPIRE"))
	configJWT := _middleware.ConfigJWT{
		SecretJWT:       os.Getenv("JWT_SECRET"),
		ExpiresDuration: int64(EXPIRE),
	}
	log.Println("jwt expire: ",EXPIRE)
	log.Println("jwt secret: ",configJWT.SecretJWT)
	e := echo.New()

	addrRepo := _driverFactory.NewAddressRepository(db)
	geoRepo := _driverFactory.NewGeolocationRepository()

	userRepo := _driverFactory.NewUserRepository(db)
	userService := _userService.NewUserService(userRepo, addrRepo, &configJWT)
	userCtrl := _userController.NewUserController(userService)
	
	laundroRepo := _driverFactory.NewLaundromatRepository(db)
	laundroService := _laundroService.NewLaundromatService(laundroRepo, addrRepo, geoRepo)
	laundroCtrl := _laundroController.NewLaundromatController(laundroService)

	productRepo := _driverFactory.NewProductRepository(db)
	productService := _productService.NewProductService(productRepo)
	productCtrl := _productController.NewProductController(productService)

	orderRepo := _driverFactory.NewOrderRepository(db)
	orderService := _orderService.NewOrderService(orderRepo, productRepo, laundroRepo)
	orderCtrl := _orderController.NewOrderController(orderService)

	routesInit := _routes.ControllerList{
		JWTMiddleware:        configJWT.Init(),
		UserController:       *userCtrl,
		LaundromatController: *laundroCtrl,
		ProductController:    *productCtrl,
		OrderController:	  *orderCtrl,
	}
	routesInit.RouteRegister(e)
	middleware.Logger(e)
	
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	log.Fatal(e.Start(":"+port))
}