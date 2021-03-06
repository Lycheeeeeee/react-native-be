package drink

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/Lycheeeeeee/react-native-be/service"
)

func GetAll(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		return s.DrinkService.GetAll(ctx)
	}
}
