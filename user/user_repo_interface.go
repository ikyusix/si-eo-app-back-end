package user

import "sieo_app/models"

type RepoInterface interface {
	FindAll() ([]*models.User, error)
	FindById(id int) (*models.User, error)
	SignUp(req *models.User) (*models.User, error)
	Update(id int, req *models.User) (*models.User, error)
	Delete(id int) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	UploadImage(id int, path string) (*models.User, error)
}
