package addresses

type Domain struct {
	ID         uint   `json:"id" gorm:"primaryKey"`
	Street     string `json:"street"`
	PostalCode int    `json:"postal_code"`
	City       string `json:"city"`
	Province   string `json:"province"`
}