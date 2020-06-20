package models

type Event struct {
	ID          int   `gorm:"primary_key"`
	EoID        int    `gorm:"colum:eo_id; not null; type:int REFERENCES eo_tb(id)"`
	Name        string `gorm:"column:event_name; not null"`
	Date        string `gorm:"column:event_date;type:date; not null"`
	Location    string `gorm:"column:event_location; not null"`
	Prince      string `gorm:"column:event_prince; not null"`
	Capacity    string `gorm:"column:event_capacity; not null"`
	Description string `gorm:"column:event_description; not null"`
	Banner      string `gorm:"column:event_banner; not null"`
}

func (e Event) TableName() string {
	return "event_tb"
}
