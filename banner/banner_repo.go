package banner

import (
	"sieo_app/models"

	"github.com/jinzhu/gorm"
)

type BannerRepo interface {
	ViewAll() (*[]models.Banner, error)
	AddPath(path []string) ([]string, error)
	Insert(banner *models.Banner, tx *gorm.DB) (*models.Banner, error)
	Delete(banner *models.Banner) error
	ViewByIdEvent(id int) (*[]models.Banner, error)
}
