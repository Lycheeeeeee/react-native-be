package account

import (
	"context"

	"github.com/Lycheeeeeee/react-native-be/errorer"

	"github.com/go-kit/kit/endpoint"
	"github.com/Lycheeeeeee/react-native-be/domain"
	"github.com/Lycheeeeeee/react-native-be/service"
)

const userType = "user"

type StatusResponse struct {
	Status string `json:"status"`
}

func Create(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(domain.Account)
		req.Type = userType

		err := s.AccountService.Create(ctx, &req)
		if err != nil {
			return nil, err
		}
		return StatusResponse{Status: "sign up success"}, nil
	}
}

type UserTypeResponse struct {
	Type string `json:"type"`
}

func Login(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(domain.Account)
		t, err := s.AccountService.GetByUserName(ctx, &req)
		if err != nil {
			return nil, errorer.ErrInvalidAccount
		}

		return UserTypeResponse{Type: t}, nil
	}
}
