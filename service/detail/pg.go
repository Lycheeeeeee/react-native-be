package detail

import (
	"context"

	"github.com/jinzhu/gorm"
	"github.com/minhkhiemm/example-go/domain"
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
