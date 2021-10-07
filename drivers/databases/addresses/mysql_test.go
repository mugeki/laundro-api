package addresses_test

import (
	"fmt"
	"testing"

	addrBusiness "laundro-api-ca/business/addresses"
	"laundro-api-ca/drivers/databases/addresses"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var addr = addrBusiness.Domain{
	ID         : 1,
	Street     : "Test Street",
	PostalCode : 12345,
	City       : "Test City",
	Province   : "Test Province",
}

func TestInsert(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	gdb, _ := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	addrRepo := addresses.NewMySQLRepository(gdb)
	
	
	mock.ExpectQuery("SELECT * FROM `addresses` WHERE (street = ? AND postal_code = ? AND city = ? AND province = ?) AND `addresses`.`id` = ? ORDER BY `addresses`.`id` LIMIT 1").
			WithArgs(addr.Street, addr.PostalCode, addr.City, addr.Province, addr.ID).
			WillReturnError(fmt.Errorf("an error"))
	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `addresses` (`street`,`postal_code`,`city`,`province`,`id`) VALUES (?,?,?,?,?)").
			WithArgs(addr.Street, addr.PostalCode, addr.City, addr.Province, addr.ID).
			WillReturnResult(sqlmock.NewResult(1, 1))
	
	mock.ExpectCommit()

	_, err = addrRepo.Insert(&addr)
	require.NoError(t, err)
}

func TestFindByID(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	gdb, _ := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	addrRepo := addresses.NewMySQLRepository(gdb)
	
	mock.ExpectQuery("SELECT * FROM `addresses` WHERE `addresses`.`id` = ? ORDER BY `addresses`.`id` LIMIT 1").
			WithArgs(1).
			WillReturnRows(
				sqlmock.NewRows([]string{"id","street","postal_code","city","province"}).
						AddRow(addr.ID, addr.Street, addr.PostalCode, addr.City, addr.Province))

	_, err = addrRepo.FindByID(addr.ID)
	require.NoError(t, err)
}

func TestFindByCity(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	gdb, _ := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	addrRepo := addresses.NewMySQLRepository(gdb)

	mock.ExpectQuery("SELECT * FROM `addresses` WHERE city = ?").
			WithArgs(addr.City).
			WillReturnRows(
				sqlmock.NewRows([]string{"id","street","postal_code","city","province"}).
						AddRow(addr.ID, addr.Street, addr.PostalCode, addr.City, addr.Province))

	_, err = addrRepo.FindByCity(addr.City)
	require.NoError(t, err)
}