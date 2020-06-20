package repo

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"

	"sieo_app/models"
	"sieo_app/user"
)

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func CreateUserRepoImlp(DB *gorm.DB) user.RepoInterface {
	return &UserRepositoryImpl{DB}
}

func (u UserRepositoryImpl) FindAll() ([]*models.User, error) {
	userList := make([]*models.User, 0)
	u.DB.Find(&userList)
	if err := u.DB.Table("user_tb").Find(&userList).Error; err != nil {
		logrus.Error(err)
		return nil, fmt.Errorf("[UserRepositoryImpl.FindAll] Get user list data: %w", err)
	}
	return userList, nil
}
func (u UserRepositoryImpl) FindById(id int) (*models.User, error) {
	user := new(models.User)
	if err := u.DB.Table("user_tb").Where("id = ?", id).First(&user).Error; err != nil {
		logrus.Error(err)
		return nil, fmt.Errorf("[UserRepositoryImpl.FindById] Get user data: %w", err)
	}
	return user, nil
}
func (u UserRepositoryImpl) SignUp(req *models.User) (*models.User, error) {
	if err := u.DB.Table("user_tb").Save(&req).Error; err != nil {
		logrus.Error(err)
		return nil, fmt.Errorf("[UserRepositoryImpl.SignUp] Add user data: %w", err)
	}
	return req, nil
}
func (u UserRepositoryImpl) Update(id int, req *models.User) (*models.User, error) {
	user := new(models.User)
	if err := u.DB.Table("user_tb").Where("id = ?", id).First(&user).Update(&req).Error; err != nil {
		logrus.Error(err)
		return nil, fmt.Errorf("[UserRepositoryImpl.Update] Update user data: %w", err)
	}
	return user, nil
}
func (u UserRepositoryImpl) Delete(id int) (*models.User, error) {
	if err := u.DB.Table("user_tb").Where("id = ?", id).Delete(&models.User{}).Error; err != nil {
		logrus.Error(err)
		return nil, fmt.Errorf("[UserRepositoryImpl.Delete] Delete user data: %w", err)
	}
	return nil, nil
}
func (u UserRepositoryImpl) GetUserByEmail(email string) (*models.User, error) {
	dataUser := new(models.User)

	if err := u.DB.Table("user_tb").Where("user_email = ?", email).First(&dataUser).Error; err != nil {
		logrus.Error(err)
		return nil, fmt.Errorf("[UserRepositoryImpl.GetUserByEmail] Get user data by email: %w", err)
	}
	return dataUser, nil
}

func (u UserRepositoryImpl) UploadImage(id int, path string) (*models.User, error) {
	dataUser := models.User{}
	err := u.DB.Model(&dataUser).Where("id = ?", id).Update("user_urlimage", path).Error
	if err != nil {
		return nil, fmt.Errorf("[UserRepositoryImpl.UploadImage] Error update model data user: %w", err)
	}
	return &dataUser, nil
}
