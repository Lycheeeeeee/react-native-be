package account

import (
	"context"

	"github.com/minhkhiemm/example-go/errorer"

	"github.com/go-kit/kit/endpoint"
	"github.com/minhkhiemm/example-go/domain"
	"github.com/minhkhiemm/example-go/service"
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
