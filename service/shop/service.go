package shop

import (
	"context"

	"github.com/Lycheeeeeee/react-native-be/domain"
)

type Service interface {
	Gets(_ context.Context) ([]domain.Shop, error)
}
