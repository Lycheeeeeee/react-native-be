package order

import "github.com/minhkhiemm/example-go/domain"

type GetRequest struct {
	ID domain.UUID
}

type MonthRequest struct {
	Month int
}
