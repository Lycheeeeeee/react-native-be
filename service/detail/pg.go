package detail

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

func (s *pgService) Create(_ context.Context, detail *domain.Detail) (*domain.Detail, error) {
	return detail, s.db.Create(&detail).Error
}

func (s *pgService) Get(_ context.Context, orderID domain.UUID) ([]domain.Detail, error) {
	details := []domain.Detail{}
	return details, s.db.Preload("Drink").Where("order_id = ?", orderID).Find(&details).Error
}
