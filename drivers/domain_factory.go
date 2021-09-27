package drivers

import (
	userDomain "laundro-api-ca/business/users"
	userDB "laundro-api-ca/drivers/databases/users"

	"gorm.io/gorm"
)


func NewUserRepository(conn *gorm.DB) userDomain.Repository{
	return userDB.NewMySQLRepository(conn)
}