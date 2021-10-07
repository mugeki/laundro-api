package request

type Addresses struct {
	Street     string `json:"street" valid:"-"`
	PostalCode int    `json:"postal_code" valid:"length(5|5)"`
	City       string `json:"city" valid:"-"`
	Province   string `json:"province" valid:"-"`
}