package repositories

import (
	"github.com/GuilhermePT1/api-social-meli/internal/domain/models"
	"gorm.io/gorm"
)

type FollowRepository struct {
	DB *gorm.DB
}

func (r *FollowRepository) Create(f *models.Follow) error {
	return r.DB.Create(f).Error
}
