package laundromats_test

import (
	"testing"
	"time"

	laundroBusiness "laundro-api-ca/business/laundromats"
	"laundro-api-ca/drivers/databases/laundromats"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var laundro = []laundroBusiness.Domain{
	{
		Id        : 1,
		Name      : "Test Laundry 1",
		OwnerID   : 1,
		AddressID : 1,
		Status    : true,
		CreatedAt : time.Now(),
		UpdatedAt : time.Now(),
	},
	{
		Id        : 2,
		Name      : "Test Laundry 2",
		OwnerID   : 2,
		AddressID : 2,
		Status    : true,
		CreatedAt : time.Now(),
		UpdatedAt : time.Now(),
	},
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
	laundroRepo := laundromats.NewMySQLRepository(gdb)
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `laundromats` (`created_at`,`updated_at`,`deleted_at`,`name`,`owner_id`,`address_id`,`status`,`id`) VALUES (?,?,?,?,?,?,?,?)").
			WithArgs(laundro[0].CreatedAt, laundro[0].UpdatedAt, nil, laundro[0].Name, laundro[0].OwnerID, laundro[0].AddressID, laundro[0].Status, laundro[0].Id).
			WillReturnResult(sqlmock.NewResult(1,1))
	mock.ExpectCommit()

	_, err = laundroRepo.Insert(&laundro[0])
	require.NoError(t, err)
}

func TestGetByAddress(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	gdb, _ := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	laundroRepo := laundromats.NewMySQLRepository(gdb)
	defer db.Close()

	input := []uint{1,2}

	mock.ExpectQuery("SELECT * FROM `laundromats` WHERE address_id IN (?,?) AND `laundromats`.`deleted_at` IS NULL").
			WithArgs(input[0],input[1]).
			WillReturnRows(
				sqlmock.NewRows([]string{"id","created_at","updated_at","name","owner_id","address_id","status"}).
						AddRow(laundro[0].Id, laundro[0].CreatedAt, laundro[0].UpdatedAt, laundro[0].Name, laundro[0].OwnerID, laundro[0].AddressID, laundro[0].Status).
						AddRow(laundro[1].Id, laundro[1].CreatedAt, laundro[1].UpdatedAt, laundro[1].Name, laundro[1].OwnerID, laundro[1].AddressID, laundro[1].Status))

	_, err = laundroRepo.GetByAddress(input)
	require.NoError(t, err)
}

func TestGetByName(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	gdb, _ := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	laundroRepo := laundromats.NewMySQLRepository(gdb)
	defer db.Close()

	input := "Laund"

	mock.ExpectQuery("SELECT * FROM `laundromats` WHERE name LIKE ? AND `laundromats`.`deleted_at` IS NULL").
			WithArgs("%"+input+"%").
			WillReturnRows(
				sqlmock.NewRows([]string{"id","created_at","updated_at","name","owner_id","address_id","status"}).
						AddRow(laundro[0].Id, laundro[0].CreatedAt, laundro[0].UpdatedAt, laundro[0].Name, laundro[0].OwnerID, laundro[0].AddressID, laundro[0].Status).
						AddRow(laundro[1].Id, laundro[1].CreatedAt, laundro[1].UpdatedAt, laundro[1].Name, laundro[1].OwnerID, laundro[1].AddressID, laundro[1].Status))

	_, err = laundroRepo.GetByName(input)
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
	laundroRepo := laundromats.NewMySQLRepository(gdb)
	defer db.Close()

	input := uint(2)

	mock.ExpectQuery("SELECT * FROM `laundromats` WHERE id = ? AND `laundromats`.`deleted_at` IS NULL ORDER BY `laundromats`.`id` LIMIT 1").
			WithArgs(input).
			WillReturnRows(
				sqlmock.NewRows([]string{"id","created_at","updated_at","name","owner_id","address_id","status"}).
						AddRow(laundro[1].Id, laundro[1].CreatedAt, laundro[1].UpdatedAt, laundro[1].Name, laundro[1].OwnerID, laundro[1].AddressID, laundro[1].Status))

	_, err = laundroRepo.GetByID(input)
	require.NoError(t, err)
}

func TestGetStatusByID(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	gdb, _ := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	laundroRepo := laundromats.NewMySQLRepository(gdb)
	defer db.Close()

	input := uint(2)

	mock.ExpectQuery("SELECT * FROM `laundromats` WHERE `laundromats`.`id` = ? AND `laundromats`.`deleted_at` IS NULL ORDER BY `laundromats`.`id` LIMIT 1").
			WithArgs(input).
			WillReturnRows(
				sqlmock.NewRows([]string{"status"}).
						AddRow(laundro[1].Status))

	resp := laundroRepo.GetStatusByID(input)
	require.Equal(t, true, resp)
}

// func TestUpdate(t *testing.T) { // masi salah
// 	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
// 	if err != nil {
// 		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
// 	}
// 	gdb, _ := gorm.Open(mysql.New(mysql.Config{
// 		Conn: db,
// 		SkipInitializeWithVersion: true,
// 	}), &gorm.Config{})
// 	laundroRepo := laundromats.NewMySQLRepository(gdb)
// 	defer db.Close()
	
// 	mock.ExpectQuery("SELECT * FROM `laundromats` WHERE id = ? AND `laundromats`.`deleted_at` IS NULL AND `laundromats`.`id` = ? ORDER BY `laundromats`.`id` LIMIT 1").
// 			WithArgs(laundro[0].Id, laundro[0].Id).
// 			WillReturnRows(
// 				sqlmock.NewRows([]string{"id","created_at","updated_at","name","owner_id","address_id","status"}).
// 						AddRow(laundro[0].Id, laundro[0].CreatedAt, laundro[0].UpdatedAt, laundro[0].Name, laundro[0].OwnerID, laundro[0].AddressID, laundro[0].Status))
// 	mock.ExpectBegin()
// 	mock.ExpectExec("UPDATE `laundromats` SET `id`=?,`created_at`=?,`updated_at`=?,`name`=?,`address_id`=?,`status`=? WHERE id = ? AND `laundromats`.`deleted_at` IS NULL AND `laundromats`.`id` = ? AND `id` = ? ORDER BY `laundromats`.`id` LIMIT 1").
// 			WithArgs(laundro[0].Id, laundro[0].CreatedAt, laundro[0].UpdatedAt, laundro[0].Name, laundro[0].AddressID, laundro[0].Status, laundro[0].Id, laundro[0].Id, laundro[0].Id).
// 			WillReturnResult(sqlmock.NewResult(1,1))
// 	mock.ExpectCommit()

// 	_, err = laundroRepo.Update(laundro[0].Id, &laundro[0])
// 	require.NoError(t, err)
// }

// func TestDelete(t *testing.T){
// 	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
// 	if err != nil {
// 		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
// 	}
// 	gdb, _ := gorm.Open(mysql.New(mysql.Config{
// 		Conn: db,
// 		SkipInitializeWithVersion: true,
// 	}), &gorm.Config{})
// 	laundroRepo := laundromats.NewMySQLRepository(gdb)
// 	defer db.Close()

// 	mock.ExpectBegin()
// 	mock.ExpectExec("DELETE FROM `domains` WHERE id = ?").
// 			WithArgs(laundro[0].Id).
// 			WillReturnResult(sqlmock.NewResult(1,1))
// 	mock.ExpectCommit()

// 	_, err = laundroRepo.Delete(laundro[0].Id)
// 	require.NoError(t, err)
// }