package event

import (
	"sieo_app/models"

	"github.com/jinzhu/gorm"
)

type EventRepo interface {
	ViewAll() (*[]models.Event, error)
	Insert(event *models.Event, tx *gorm.DB) (*models.Event, error)
	ViewById(id int) (*models.Event, error)
	BeginTrans() *gorm.DB
	Delete(id int) error
	ViewListEvent() (*[]models.ListEvent, error)
}
