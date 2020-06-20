package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"sieo_app/event"
	"sieo_app/models"
	"strings"

	"github.com/jinzhu/gorm"
)

type EventRepoImpl struct {
	DB  *gorm.DB
	qry *sql.DB
}

func CreateEventRepoImpl(DB *gorm.DB, db *sql.DB) event.EventRepo {
	return &EventRepoImpl{DB, db}
}

func (e *EventRepoImpl) BeginTrans() *gorm.DB {
	return e.DB.Begin()
}

func (e *EventRepoImpl) ViewAll() (*[]models.Event, error) {
	var event []models.Event
	err := e.DB.Find(&event).Error
	if err != nil {
		return nil, fmt.Errorf("[EventRepoImpl.ViewAll] Error When Execute Query")
	}
	return &event, nil
}

func (e *EventRepoImpl) ViewById(id int) (*models.Event, error) {
	event := new(models.Event)

	if err := e.DB.Table("event_tb").Where("ID = ?", id).First(&event).Error; err != nil {
		fmt.Errorf("[EventRepoImpl.ViewById] When Execute Query")
		return nil, errors.New("get event data: error")
	}
	return event, nil
}

func (e *EventRepoImpl) Insert(event *models.Event, tx *gorm.DB) (*models.Event, error) {
	err := tx.Save(event).Error
	if err != nil {
		return nil, fmt.Errorf("[EventRepoImpl.Insert] Error insert data")
	}
	return event, nil
}

func (e *EventRepoImpl) ViewListEvent() (*[]models.ListEvent, error) {
	var listEvents []models.ListEvent
	var listEvent models.ListEvent
	Query := "select distinct on (event_tb.event_name) event_tb.id, event_tb.event_name, event_tb.event_date, event_tb.event_location, event_tb.event_capacity, event_tb.event_prince, event_tb.event_description, banner_tb.banner_url from event_tb inner join banner_tb on (event_tb.id = banner_tb.event_id)"
	rows, err := e.qry.Query(Query)
	if err != nil {
		fmt.Println("[EventRepoMysqlImpl.Insert] Error insert data")
		return nil, fmt.Errorf("[EventRepoMysqlImpl.Insert] Error insert data")
	}

	for rows.Next() {
		var err = rows.Scan(&listEvent.Id, &listEvent.Name, &listEvent.Date, &listEvent.Location, &listEvent.Capacity, &listEvent.Prince, &listEvent.Description, &listEvent.UrlBanner)

		text := strings.Replace(listEvent.Date, "T00:00:00Z", "", 1)
		listEvent.Date = text

		if err != nil {
			fmt.Errorf("[EventRepoImpl.ViewList] Error When Scan rows next : %v ", err)
			return nil, fmt.Errorf("Oooops, there's something wrong")
		}
		listEvents = append(listEvents, listEvent)
	}

	return &listEvents, nil
}

func (e *EventRepoImpl) Delete(id int) error {
	fmt.Println(id)
	err := e.DB.Where("id = ?", id).Delete(&models.Event{}).Error
	if err != nil {
		return fmt.Errorf("[EventRepoImpl.Delete] Failed Delete Data")
	}
	return nil
}
