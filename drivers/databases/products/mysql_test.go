package products_test

import (
	productBusiness "laundro-api-ca/business/products"
	"laundro-api-ca/drivers/databases/products"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var product = productBusiness.Domain{
    Id             : 1,
    KgLimit        : 1,
    KgPrice        : 1000,
    EstimatedHour  : 1,
    CategoryID     : 1,
    CategoryName   : "Test Category",
    LaundromatID   : 1,
    LaundromatName : "Test Laundy",
    CreatedAt      : time.Now(),
    UpdatedAt      : time.Now(),
}

func TestInsert(t *testing.T){
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	gdb, _ := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	productRepo := products.NewMySQLRepository(gdb)
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `categories` (`name`,`id`) VALUES (?,?) ON DUPLICATE KEY UPDATE `id`=`id`").
		WithArgs(product.CategoryName, product.CategoryID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec("INSERT INTO `products` (`created_at`,`updated_at`,`deleted_at`,`kg_limit`,`kg_price`,`estimated_hour`,`category_id`,`laundromat_id`,`id`) VALUES (?,?,?,?,?,?,?,?,?)").
		WithArgs(product.CreatedAt, product.UpdatedAt, nil, product.KgLimit, product.KgPrice, product.EstimatedHour, product.CategoryID, product.LaundromatID, product.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	_, err = productRepo.Insert(&product)
	require.NoError(t, err)
}

func TestGetAllByLaundromat(t *testing.T){
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	gdb, _ := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	productRepo := products.NewMySQLRepository(gdb)
	defer db.Close()
	
	mock.ExpectQuery("SELECT `products`.`id`,`products`.`created_at`,`products`.`updated_at`,`products`.`deleted_at`,`products`.`kg_limit`,`products`.`kg_price`,`products`.`estimated_hour`,`products`.`category_id`,`products`.`laundromat_id`,`Category`.`id` AS `Category__id`,`Category`.`name` AS `Category__name` FROM `products` LEFT JOIN `categories` `Category` ON `products`.`category_id` = `Category`.`id` WHERE laundromat_id = ? AND `products`.`deleted_at` IS NULL").
			WithArgs(product.LaundromatID).
			WillReturnRows(
				sqlmock.NewRows([]string{"id","created_at","updated_at","deleted_at","kg_limit","kg_price","estimated_hour","category_id","laundromat_id"}).
						AddRow(product.Id,product.CreatedAt,product.UpdatedAt,nil,product.KgLimit,product.KgPrice,product.EstimatedHour,product.CategoryID,product.LaundromatID))

	_, err = productRepo.GetAllByLaundromat(product.LaundromatID)
	require.NoError(t, err)
}

// func TestDelete(t *testing.T){
// 	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
// 	if err != nil {
// 		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
// 	}
// 	gdb, _ := gorm.Open(mysql.New(mysql.Config{
// 		Conn:                      db,
// 		SkipInitializeWithVersion: true,
// 	}), &gorm.Config{})
// 	productRepo := products.NewMySQLRepository(gdb)
// 	defer db.Close()

// 	mock.ExpectBegin()
// 	mock.ExpectExec("UPDATE `products` SET `deleted_at`=? WHERE id = ? AND `products`.`deleted_at` IS NULL").
// 			WithArgs(time.Now(), product.Id).
// 			WillReturnResult(sqlmock.NewResult(1,1))
// 	mock.ExpectCommit()

// 	_, err = productRepo.Delete(product.Id)
// 	require.NoError(t, err)
// }