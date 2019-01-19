package account

import (
	"context"

	"github.com/Lycheeeeeee/react-native-be/domain"
)

type Service interface {
	Create(ctx context.Context, account *domain.Account) error
	GetByUserName(ctx context.Context, account *domain.Account) (string, domain.UUID, error)
}
