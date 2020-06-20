package models

type Banner struct {
	ID        int    `gorm:"primary_key"`
	EventId   int    `gorm:"colum:event_id; not null; type:int REFERENCES event_tb(id)"`
	UrlBanner string `gorm:"column:banner_url; not null"`
}

func (e Banner) TableName() string {
	return "banner_tb"
}
