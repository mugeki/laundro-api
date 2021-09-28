package drivers

import (
	addressDomain "laundro-api-ca/business/addresses"
	addressDB "laundro-api-ca/drivers/databases/addresses"

	userDomain "laundro-api-ca/business/users"
	userDB "laundro-api-ca/drivers/databases/users"

	"gorm.io/gorm"
)


func NewUserRepository(conn *gorm.DB) userDomain.Repository{
	return userDB.NewMySQLRepository(conn)
}

func NewAddressRepository(conn *gorm.DB) addressDomain.Repository{
	return addressDB.NewMySQLRepository(conn)
}