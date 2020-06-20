package models

type User struct {
	ID        int    `gorm:"primary_key"`
	Fullname  string `gorm:"column:user_fullname"`
	Gender    string `gorm:"column:user_gender"`
	BirthDate string `gorm:"column:user_birthdate"`
	Email     string `gorm:"column:user_email"`
	Username  string `gorm:"column:user_username"`
	Password  string `gorm:"column:user_password"`
	UrlImage  string `gorm:"column:user_urlimage"`
	Status    string `gorm:"column:user_status"`
}

func (h User) TableName() string {
	return "user_tb"
}
