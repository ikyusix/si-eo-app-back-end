package event

import (
	"sieo_app/models"
)

type EventService interface {
	ViewAll() (*[]models.Event, error)
	InsertEvents(event *models.Event, path []string) (*models.Event, error)
	ViewById(id int) (*models.Event, error)
	Delete(id int) error
	ViewListEvent() (*[]models.ListEvent, error)
}
