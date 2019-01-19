package detail

import (
	"context"

	"github.com/Lycheeeeeee/react-native-be/domain"
)

type Service interface {
	Create(ctx context.Context, detail *domain.Detail) (*domain.Detail, error)
	Get(ctx context.Context, orderID domain.UUID) ([]domain.Detail, error)
}
