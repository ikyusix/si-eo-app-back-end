package service

import (
	"fmt"
	"sieo_app/banner"
	"sieo_app/event"
	"sieo_app/models"
)

type EventServiceImpl struct {
	eventRepo  event.EventRepo
	bannerRepo banner.BannerRepo
}

func CreateEventService(eventRepo event.EventRepo, bannerRepo banner.BannerRepo) event.EventService {
	return &EventServiceImpl{eventRepo, bannerRepo}
}

func (e *EventServiceImpl) ViewAll() (*[]models.Event, error) {
	return e.eventRepo.ViewAll()
}

func (e *EventServiceImpl) Delete(id int) error {
	return e.eventRepo.Delete(id)
}

func (e *EventServiceImpl) ViewListEvent() (*[]models.ListEvent, error) {
	return e.eventRepo.ViewListEvent()
}

func (e *EventServiceImpl) InsertEvents(event *models.Event, path []string) (*models.Event, error) {
	tx := e.eventRepo.BeginTrans()

	events, err := e.eventRepo.Insert(event, tx)
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("[EventServiceImpl.InsertEvents] Error insert table event data")
	}

	for i := 0; i < len(path); i++ {
		var banner = models.Banner{
			EventId:   events.ID,
			UrlBanner: path[i],
		}

		_, err := e.bannerRepo.Insert(&banner, tx)
		if err != nil {
			tx.Rollback()
			return nil, fmt.Errorf("[EventRepoMysqlImpl.Insert] Error insert table image data")
		}
	}
	tx.Commit()
	return events, nil
}

func (e *EventServiceImpl) ViewById(id int) (*models.Event, error) {
	return e.eventRepo.ViewById(id)
}

// func (e *EventServiceImpl) InsertAll(event *models.Event) (*models.Event, error) {

// 	event, err := e.eventRepo.Insert(event)
// 	AddPath(path []string) ([]string, error)

// }
