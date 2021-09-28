package geolocation

type Domain struct {
	IP   string
	City string
}

type Repository interface {
	GetLocationByIP() (Domain, error)
}