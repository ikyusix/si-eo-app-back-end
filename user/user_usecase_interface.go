package user

import "sieo_app/models"

type UseCaseInterface interface {
	FindAll() (map[string]interface{}, error)
	FindById(id int) (map[string]interface{}, error)
	SignUp(req *models.User) (map[string]interface{}, error)
	Update(id int, req *models.User) (map[string]interface{}, error)
	Delete(id int) (map[string]interface{}, error)
	GetUserByEmail(email string) (*models.User, error)
	UploadImage(id int, path string) (*models.User, string, error)
}
