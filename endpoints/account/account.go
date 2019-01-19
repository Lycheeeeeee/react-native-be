package account

import (
	"context"

	"github.com/Lycheeeeeee/react-native-be/errorer"

	"github.com/Lycheeeeeee/react-native-be/domain"
	"github.com/Lycheeeeeee/react-native-be/service"
	"github.com/go-kit/kit/endpoint"
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

type LoginStaffResponse struct {
	Type   string      `json:"type"`
	ShopID domain.UUID `json:"shop_id"`
}

func Login(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(domain.Account)
		t, s, err := s.AccountService.GetByUserName(ctx, &req)
		if err != nil {
			return nil, errorer.ErrInvalidAccount
		}

		if t == "staff" {
			return LoginStaffResponse{
				Type:   t,
				ShopID: s,
			}, nil

		}

		return UserTypeResponse{Type: t}, nil
	}
}
