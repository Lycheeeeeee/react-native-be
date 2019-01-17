package order

import "github.com/Lycheeeeeee/react-native-be/domain"

type GetRequest struct {
	ID domain.UUID
}

type MonthRequest struct {
	Month int
}
