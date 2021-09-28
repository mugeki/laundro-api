package users

import (
	"fmt"
	"laundro-api-ca/app/middleware"
	"laundro-api-ca/business"
	"laundro-api-ca/business/addresses"
	"laundro-api-ca/helper/encrypt"
)

type userService struct {
	userRepository 		Repository
	addrRepository		addresses.Repository
	jwtAuth		   		*middleware.ConfigJWT
}

func NewUserService(userRepo Repository, addrRepo addresses.Repository, jwtauth *middleware.ConfigJWT) Service {
	return &userService{
		userRepository: userRepo,
		addrRepository: addrRepo,
		jwtAuth: jwtauth,
	}
}

func (service *userService) Register(userData *Domain, addressData *addresses.Domain) (Domain, error){
	newAddr, err := service.addrRepository.Insert(addressData)
	userData.AddressID = newAddr.ID
	fmt.Println(userData)
	res, err := service.userRepository.Register(userData)
	if res == (Domain{}) {
		return Domain{}, business.ErrDuplicateData
	}
	if err != nil {
		return Domain{}, err
	}
	return res, nil
}

func (service *userService) Login(username, password string) (string, error){
	userDomain, err := service.userRepository.GetByUsername(username)
	if err != nil {
		return "", business.ErrInvalidLoginInfo
	}

	if !encrypt.ValidateHash(password, userDomain.Password){
		return "", business.ErrInvalidLoginInfo
	}

	token := service.jwtAuth.GenerateToken(int(userDomain.Id))
	return token, nil
}

func (service *userService) GetByID(id uint) (Domain, error){
	userDomain, err := service.userRepository.GetByID(id)
	if err != nil {
		return Domain{}, err
	}
	return userDomain, nil
}