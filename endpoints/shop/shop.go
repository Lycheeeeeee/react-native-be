package shop

import (
	"context"

	"github.com/Lycheeeeeee/react-native-be/service"
	"github.com/go-kit/kit/endpoint"
)

func GetAll(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		shop, err := s.ShopService.Gets(ctx)
		if err != nil {
			return nil, err
		}

		return shop, nil
	}
}
