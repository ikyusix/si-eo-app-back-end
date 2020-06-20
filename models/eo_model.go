package models

type Eo struct {
	ID          int    `gorm:"primary_key"`
	UserID      int    `gorm:"colum:user_id; not null; type:int REFERENCES user_tb(id)"`
	Name        string `gorm:"column:eo_name" ;json:"name" validate:"required"`
	Identity    string `gorm:"column:eo_identity" ;json:"identity" validate:"required"`
	ImgIdentity string `gorm:"column:eo_identity_img" ;json:"imgIdentity" validate:"required"`
	License     string `gorm:"column:eo_license" ;json:"license" validate:"required"`
	ImgLicense  string `gorm:"column:eo_license_img" ;json:"imgLicense" validate:"required"`
	Address     string `gorm:"column:eo_address" ;json:"address" validate:"required"`
	Phone       string `gorm:"column:eo_numberphone" ;json:"phone" validate:"required"`
	Website     string `gorm:"column:eo_website" ;json:"website" validate:"required"`
	Instagram   string `gorm:"column:eo_instagram" ;json:"instagram" validate:"required"`
	Facebook    string `gorm:"column:eo_facebook" ;json:"facebook" validate:"required"`
	Twitter     string `gorm:"column:eo_twitter" ;json:"twitter" validate:"required"`
	Status      string `gorm:"column:eo_status" ;json:"status" validate:"required"`
}

func (e Eo) TableName() string {
	return "eo_tb"
}
