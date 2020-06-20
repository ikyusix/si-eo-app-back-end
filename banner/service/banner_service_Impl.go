package service

import (
	"sieo_app/banner"
	"sieo_app/models"
)

type BannerServiceImpl struct {
	bannerRepo banner.BannerRepo
}

func CreateBannerService(bannerRepo banner.BannerRepo) banner.BannerService {
	return &BannerServiceImpl{bannerRepo}
}

func (e *BannerServiceImpl) ViewAll() (*[]models.Banner, error) {
	return e.bannerRepo.ViewAll()
}

func (e *BannerServiceImpl) ViewByIdEvent(id int) (*[]models.Banner, error) {
	return e.bannerRepo.ViewByIdEvent(id)
}

// func (e *BannerServiceImpl) Insert(banner *models.Banner) (*models.Banner, error) {
// 	return e.bannerRepo.Insert(banner)
// }

func (e *BannerServiceImpl) AddPath(path []string) ([]string, error) {
	return e.bannerRepo.AddPath(path)
}

func (e *BannerServiceImpl) Delete(banner *models.Banner) error {
	return e.bannerRepo.Delete(banner)
}
