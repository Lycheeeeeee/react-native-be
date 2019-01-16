package drink

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/minhkhiemm/example-go/service"
)

func GetAll(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		return s.DrinkService.GetAll(ctx)
	}
}
