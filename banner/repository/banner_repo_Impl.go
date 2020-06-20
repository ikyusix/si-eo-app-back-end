package banner

import (
	"fmt"
	"sieo_app/banner"
	"sieo_app/models"

	"github.com/jinzhu/gorm"
)

type BannerRepoImpl struct {
	DB *gorm.DB
}

func CreateBannerRepoImpl(DB *gorm.DB) banner.BannerRepo {
	return &BannerRepoImpl{DB}
}

func (e *BannerRepoImpl) ViewAll() (*[]models.Banner, error) {
	var banner []models.Banner
	err := e.DB.Find(&banner).Error
	if err != nil {
		return nil, fmt.Errorf("[EventRepoMysqlImpl.ViewAll] Error Wwhen Execute query sql")
	}
	return &banner, nil
}

func (e *BannerRepoImpl) ViewByIdEvent(id int) (*[]models.Banner, error) {
	var banner []models.Banner
	if err := e.DB.Table("banner_tb").Where("event_id = ?", id).Find(&banner).Error; err != nil {
		fmt.Print("[BannerRepoImpl.ViewByIdEvent] When Execute Query")
		return nil, fmt.Errorf("get banner data: error")
	}
	return &banner, nil
}

func (e *BannerRepoImpl) Insert(banner *models.Banner, tx *gorm.DB) (*models.Banner, error) {
	err := tx.Save(banner).Error
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("[BannerRepoMysqlImpl.Insert] Error insert data")
	}
	return banner, nil
}

func (e *BannerRepoImpl) Delete(banner *models.Banner) error {
	err := e.DB.Delete(&banner)
	if err != nil {
		return fmt.Errorf("[BannerRepoImpl.Insert] failed delete data")
	}
	return nil
}

func (e *BannerRepoImpl) AddPath(path []string) ([]string, error) {
	fmt.Println(path)
	return path, nil
}
