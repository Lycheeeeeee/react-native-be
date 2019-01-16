package drink

import (
	"context"

	"github.com/minhkhiemm/example-go/domain"
)

type Service interface {
	GetAll(ctx context.Context) ([]domain.Drink, error)
}
