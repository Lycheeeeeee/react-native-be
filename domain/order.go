package domain

import (
	"time"
)

type Order struct {
	Model
	OrderTime   time.Time `json:"order_time"`
	ReceiveTime *int      `json:"receive_time"`
	Status      *string   `json:"status"`
	ShopID      UUID      `json:"shop_id"`
	DetailID    UUID      `json:"detail_id"`
	Details     []Detail  `json:"details"`
	AccountID   UUID      `json:"account_id"`
	TotalPrice  int       `json:"total_price"`
}

type OrderDate struct {
	Date   int
	Month  int
	Year   int
	ShopID UUID
}
