package banner

import "sieo_app/models"

type BannerService interface {
	ViewAll() (*[]models.Banner, error)
	AddPath(path []string) ([]string, error)
	Delete(banner *models.Banner) error
	ViewByIdEvent(id int) (*[]models.Banner, error)
	// Insert(banner *models.Banner) (*models.Banner, error)
}
