package domain

type Detail struct {
	Model
	Quantity int   `json:"quantity"`
	DrinkID  UUID  `json:"drink_id"`
	Drink    Drink `json:"drink"`
	OrderID  UUID  `json:"order_id"`
}
