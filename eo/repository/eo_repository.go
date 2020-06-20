package repository

import (
	"errors"
	"sieo_app/eo"
	"sieo_app/models"

	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

type EoRepositoryImpl struct {
	Conn *gorm.DB
}

func CreateEoRepoImpl(DB *gorm.DB) eo.EoRepoInterface {
	return &EoRepositoryImpl{DB}
}

func (e *EoRepositoryImpl) ReadAll() ([]*models.Eo, error) {
	eoList := make([]*models.Eo, 0)

	if err := e.Conn.Table("eo_tb").Find(&eoList).Error; err != nil {
		logrus.Error(err)
		return nil, errors.New("Get EO List Data: Error")
	}

	return eoList, nil
}

func (e *EoRepositoryImpl) Insert(req *models.Eo) (*models.Eo, error) {
	if err := e.Conn.Table("eo_tb").Save(&req).Error; err != nil {
		logrus.Error(err)
		return nil, errors.New("Add EO Data : Failed To Add Data")
	}
	return req, nil
}

func (e *EoRepositoryImpl) UpdateEo(id int, req *models.Eo) (*models.Eo, error) {
	eo := new(models.Eo)

	if err := e.Conn.Table("eo_tb").Where("user_id=?", id).First(&eo).Update(&req).Error; err != nil {
		logrus.Error(err)
		return nil, errors.New("Update EO Data : Failed To Update Data")
	}
	return eo, nil
}
