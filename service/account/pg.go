package account

import (
	"context"

	"github.com/jinzhu/gorm"

	"github.com/minhkhiemm/example-go/domain"
	"github.com/minhkhiemm/example-go/errorer"
)

type pgService struct {
	db *gorm.DB
}

func NewPGService(db *gorm.DB) Service {
	return &pgService{
		db: db,
	}
}

func (s *pgService) Create(_ context.Context, account *domain.Account) error {
	err := s.db.Where("user_name = ?", account.UserName).First(&domain.Account{}).Error
	if err != nil {
		return s.db.Create(&account).Error
	}

	return errorer.ErrInvalidUserName
}

func (s *pgService) GetByUserName(_ context.Context, account *domain.Account) (string, error) {
	err := s.db.Where("deleted_at IS NULL AND user_name = ? AND digest_password = ?", account.UserName, account.DigestPassword).First(&account).Error
	if err != nil {
		return "", err
	}

	return account.Type, nil
}
