package endpoints

import (
	"github.com/go-kit/kit/endpoint"

	"github.com/minhkhiemm/example-go/endpoints/account"
	"github.com/minhkhiemm/example-go/endpoints/drink"
	"github.com/minhkhiemm/example-go/endpoints/order"
	"github.com/minhkhiemm/example-go/service"
)

// Endpoints .
type Endpoints struct {
	// Orders
	GetAllOrderByDate endpoint.Endpoint
	CreateOrder       endpoint.Endpoint
	GetOrderByID      endpoint.Endpoint
	UpdateOrder       endpoint.Endpoint
	GetOrderByShopID  endpoint.Endpoint
	GetOrderByMonth   endpoint.Endpoint
	// Accounts
	CreateAccount endpoint.Endpoint
	Login         endpoint.Endpoint
	// Drinks
	GetAllDrink endpoint.Endpoint
}

// MakeServerEndpoints returns an Endpoints struct
func MakeServerEndpoints(s service.Service) Endpoints {
	return Endpoints{
		// Orders
		GetAllOrderByDate: order.GetAllByDate(s),
		CreateOrder:       order.Create(s),
		GetOrderByID:      order.Get(s),
		UpdateOrder:       order.Update(s),
		GetOrderByShopID:  order.GetByShopID(s),
		GetOrderByMonth:   order.GetByMonth(s),
		// Accounts
		CreateAccount: account.Create(s),
		Login:         account.Login(s),
		// Drinks
		GetAllDrink: drink.GetAll(s),
	}
}
