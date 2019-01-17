package service

import (
	"github.com/Lycheeeeeee/react-native-be/service/shop"
	"github.com/Lycheeeeeee/react-native-be/service/account"
	"github.com/Lycheeeeeee/react-native-be/service/detail"
	"github.com/Lycheeeeeee/react-native-be/service/drink"
	"github.com/Lycheeeeeee/react-native-be/service/order"
)

// Service define list of all services in projects
type Service struct {
	OrderService   order.Service
	AccountService account.Service
	DetailService  detail.Service
	DrinkService   drink.Service
	ShopService    shop.Service
}
