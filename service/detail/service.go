package detail

import (
	"context"

	"github.com/minhkhiemm/example-go/domain"
)

type Service interface {
	Create(ctx context.Context, detail *domain.Detail) (*domain.Detail, error)
}
