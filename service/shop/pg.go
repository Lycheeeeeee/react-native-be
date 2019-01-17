package shop

import (
	"context"

	"github.com/Lycheeeeeee/react-native-be/domain"
	"github.com/jinzhu/gorm"
)

type pgService struct {
	db *gorm.DB
}

func NewPGService(db *gorm.DB) Service {
	return &pgService{
		db: db,
	}
}

func (s *pgService) Gets(_ context.Context) ([]domain.Shop, error) {
	shops := []domain.Shop{}
	return shops, s.db.Where("deleted_at IS NULL").Find(&shops).Error
}
