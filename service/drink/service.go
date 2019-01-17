package drink

import (
	"context"

	"github.com/Lycheeeeeee/react-native-be/domain"
)

type Service interface {
	GetAll(ctx context.Context) ([]domain.Drink, error)
}
