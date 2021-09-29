package drivers

import (
	addressDomain "laundro-api-ca/business/addresses"
	addressDB "laundro-api-ca/drivers/databases/addresses"

	userDomain "laundro-api-ca/business/users"
	userDB "laundro-api-ca/drivers/databases/users"

	productDomain "laundro-api-ca/business/products"
	productDB "laundro-api-ca/drivers/databases/products"

	laundroDomain "laundro-api-ca/business/laundromats"
	laundroDB "laundro-api-ca/drivers/databases/laundromats"

	geolocationDomain "laundro-api-ca/business/geolocation"
	ipAPI "laundro-api-ca/drivers/thirdparties/ipapi"

	"gorm.io/gorm"
)


func NewUserRepository(conn *gorm.DB) userDomain.Repository{
	return userDB.NewMySQLRepository(conn)
}

func NewAddressRepository(conn *gorm.DB) addressDomain.Repository{
	return addressDB.NewMySQLRepository(conn)
}

func NewLaundromatRepository(conn *gorm.DB) laundroDomain.Repository{
	return laundroDB.NewMySQLRepository(conn)
}

func NewGeolocationRepository() geolocationDomain.Repository{
	return ipAPI.NewIpAPI()
}

func NewProductRepository(conn *gorm.DB) productDomain.Repository{
	return productDB.NewMySQLRepository(conn)
}