package orders

import (
	"errors"
	"laundro-api-ca/business/orders"

	"gorm.io/gorm"
)

type mysqlOrdersRepository struct {
	Conn *gorm.DB
}

func NewMySQLRepository(conn *gorm.DB) orders.Repository {
	return &mysqlOrdersRepository{
		Conn: conn,
	}
}

func (mysqlRepo *mysqlOrdersRepository) Create(orderData *orders.Domain) (orders.Domain, error){
	rec := fromDomain(*orderData)
	err := mysqlRepo.Conn.Create(&rec).Error
	if err != nil {
		return orders.Domain{}, err
	}
	return rec.toDomain(), err
}

func (mysqlRepo *mysqlOrdersRepository) GetByUserID(userId uint) ([]orders.Domain, error){
	rec := []Orders{}

	err := mysqlRepo.Conn.Joins("Laundromat").Joins("Payment").Joins("Product").Find(&rec, "user_id = ?", userId).Error
	if len(rec) == 0 {
		err = errors.New("Not Found")
		return []orders.Domain{}, err
	}
	orders := toDomainArray(rec)
	return orders, nil
}

func (mysqlRepo *mysqlOrdersRepository) GetPaymentGateway(paymentId int) (string, error){
	paymentRec := Payment{}
	err := mysqlRepo.Conn.First(&paymentRec, paymentId).Error
	if err != nil {
		return "", err
	}
	return paymentRec.Gateway, nil
}