package request

type Addresses struct {
	Street     string `json:"street"`
	PostalCode int    `json:"postal_code"`
	City       string `json:"city"`
	Province   string `json:"province"`
}