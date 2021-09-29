package laundromats

import (
	"laundro-api-ca/business"
	"laundro-api-ca/business/addresses"
	"laundro-api-ca/business/geolocation"
)

type laundroService struct {
	laundroRepository 	Repository
	addrRepository		addresses.Repository
	geoRepository 		geolocation.Repository
}

func NewLaundromatService(laundroRepo Repository, addrRepo addresses.Repository, geoRepo geolocation.Repository) Service {
	return &laundroService{
		laundroRepository: laundroRepo,
		addrRepository:    addrRepo,
		geoRepository:	   geoRepo,
	}
}

func (service *laundroService) Insert(userID uint, laundroData *Domain, addressData *addresses.Domain) (Domain, error){
	newAddr, err := service.addrRepository.Insert(addressData)
	laundroData.AddressID = newAddr.ID
	laundroData.OwnerID = userID
	res, err := service.laundroRepository.Insert(laundroData)
	if res == (Domain{}) {
		return Domain{}, business.ErrDuplicateData
	}
	if err != nil {
		return Domain{}, err
	}
	return res, nil
}

func (service *laundroService) GetByIP() ([]Domain, error){
	location, err := service.geoRepository.GetLocationByIP()
	if err != nil {
		return []Domain{}, business.ErrInternalServer
	}
	addrData, err := service.addrRepository.FindByCity(location.City)
	if err != nil {
		return []Domain{}, business.ErrNearestLaundromatNotFound
	}
	
	addressID := []uint{}
	for _, val := range addrData{
		addressID = append(addressID,val.ID)
	}
	laundroDomain, err := service.laundroRepository.GetByAddress(addressID)
	if err != nil {
		return []Domain{}, business.ErrNearestLaundromatNotFound
	}
	return laundroDomain, nil
}

func (service *laundroService) GetByName(name string) ([]Domain, error){
	laundroDomain, err := service.laundroRepository.GetByName(name)
	if err != nil {
		return []Domain{}, business.ErrLaundromatNotFound
	}
	return laundroDomain, nil

}

func (service *laundroService) GetByID(id uint) (Domain, error){
	laundroDomain, err := service.laundroRepository.GetByID(id)
	if err != nil {
		return Domain{}, business.ErrLaundromatNotFound
	}
	return laundroDomain, nil
}

func (service *laundroService) Update(id uint, laundroData *Domain, addressData *addresses.Domain) (Domain, error){
	newAddr, err := service.addrRepository.Insert(addressData)
	laundroData.AddressID = newAddr.ID
	laundroDomain, err := service.laundroRepository.Update(id, laundroData)
	if err != nil {
		return Domain{}, business.ErrLaundromatNotFound
	}
	return laundroDomain, nil
}

func (service *laundroService) Delete(id uint) (string, error){
	res, err := service.laundroRepository.Delete(id)
	if err != nil {
		return "", business.ErrLaundromatNotFound
	}
	return res, nil
}