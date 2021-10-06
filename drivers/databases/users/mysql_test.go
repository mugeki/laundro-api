package users_test

import (
	userBusiness "laundro-api-ca/business/users"
	"laundro-api-ca/drivers/databases/users"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var user = userBusiness.Domain{
	Id          : 1,
    Username    : "testUser",
    Password    : "testPassword",
    Email       : "test@gmail.com",
    Fullname    : "Test John",
    DateOfBirth : time.Date(2001,time.January,1,0,0,0,0,time.Local),
    PhoneNumber : "123456789",
    RoleID      : 1,
    AddressID   : 1,
    CreatedAt   : time.Now(),
    UpdatedAt   : time.Now(),
}

func TestInsert(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	gdb, _ := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	userRepo := users.NewMySQLRepository(gdb)
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `users` (`created_at`,`updated_at`,`deleted_at`,`username`,`password`,`email`,`fullname`,`date_of_birth`,`phone_number`,`role_id`,`address_id`,`id`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?)").
		WithArgs(user.CreatedAt, user.UpdatedAt, nil, user.Username, user.Password, user.Email, user.Fullname, user.DateOfBirth, user.PhoneNumber, user.RoleID, user.AddressID, user.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	_, err = userRepo.Register(&user)
	require.NoError(t, err)
}

func TestGetByUsername(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	gdb, _ := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	userRepo := users.NewMySQLRepository(gdb)
	defer db.Close()
	
	input := "testUser"

	mock.ExpectQuery("SELECT * FROM `users` WHERE username = ? AND `users`.`deleted_at` IS NULL ORDER BY `users`.`id` LIMIT 1").
			WithArgs(input).
			WillReturnRows(
				sqlmock.NewRows([]string{"id","username","password","email","fullname","date_of_birth","phone_number","role_id","address_id","created_at","updated_at"}).
						AddRow(user.Id, user.Username, user.Password, user.Email, user.Fullname, user.DateOfBirth, user.PhoneNumber, user.RoleID, user.AddressID, user.CreatedAt, user.UpdatedAt))

	_, err = userRepo.GetByUsername(input)
	require.NoError(t, err)
}

func TestGetByID(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	gdb, _ := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	userRepo := users.NewMySQLRepository(gdb)
	defer db.Close()
	
	input := uint(1)

	mock.ExpectQuery("SELECT * FROM `users` WHERE id = ? AND `users`.`deleted_at` IS NULL ORDER BY `users`.`id` LIMIT 1").
			WithArgs(input).
			WillReturnRows(
				sqlmock.NewRows([]string{"id","username","password","email","fullname","date_of_birth","phone_number","role_id","address_id","created_at","updated_at"}).
						AddRow(user.Id, user.Username, user.Password, user.Email, user.Fullname, user.DateOfBirth, user.PhoneNumber, user.RoleID, user.AddressID, user.CreatedAt, user.UpdatedAt))

	_, err = userRepo.GetByID(input)
	require.NoError(t, err)
}