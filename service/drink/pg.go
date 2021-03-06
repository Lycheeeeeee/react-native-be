package drink

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

func (s *pgService) GetAll(_ context.Context) ([]domain.Drink, error) {
	drinks := []domain.Drink{}
	err := s.db.Debug().Table("details").Select("d.*, sum(details.quantity) as times_order").
		Joins("JOIN drinks as d ON d.id = details.drink_id").
		Group("d.id").Order("times_order desc").Find(&drinks).Error
	if err != nil {
		return nil, err
	}
	return drinks, nil
}

func (s *pgService) Get(_ context.Context, name string) (*domain.Drink, error) {
	drink := domain.Drink{}

	return &drink, s.db.Where("name = ?", name).First(&drink).Error
}
