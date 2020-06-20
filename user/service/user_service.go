package usecase

import (
	"sieo_app/common"
	"sieo_app/models"
	"sieo_app/user"
)

type UserServiceImpl struct {
	UserRepo user.RepoInterface
}

func CreateUserUsecaseImpl(userRepo user.RepoInterface) user.UseCaseInterface {
	return &UserServiceImpl{UserRepo: userRepo}
}

func (u UserServiceImpl) FindAll() (map[string]interface{}, error) {
	response, err := u.UserRepo.FindAll()
	if err != nil {
		return common.Message(false, err.Error()), err
	}
	mapResponse := common.Message(true, "Read all user data: success")
	mapResponse["response"] = response
	return mapResponse, nil
}
func (u UserServiceImpl) FindById(id int) (map[string]interface{}, error) {
	response, err := u.UserRepo.FindById(id)
	if err != nil {
		return common.Message(false, err.Error()), err
	}

	mapResponse := common.Message(true, "Read user data: success")
	mapResponse["response"] = response
	return mapResponse, nil
}
func (u UserServiceImpl) SignUp(req *models.User) (map[string]interface{}, error) {
	response, err := u.UserRepo.SignUp(req)
	if err != nil {
		return common.Message(false, err.Error()), err
	}
	mapResponse := common.Message(true, "Create user data: success")
	mapResponse["response"] = response
	return mapResponse, nil
}
func (u UserServiceImpl) Update(id int, req *models.User) (map[string]interface{}, error) {
	response, err := u.UserRepo.Update(id, req)
	if err != nil {
		return common.Message(false, err.Error()), err
	}

	mapResponse := common.Message(true, "Update user data: success")
	mapResponse["response"] = response
	return mapResponse, nil
}
func (u UserServiceImpl) Delete(id int) (map[string]interface{}, error) {
	_, err := u.UserRepo.Delete(id)
	if err != nil {
		return common.Message(false, err.Error()), err
	}

	mapResponse := common.Message(true, "Delete user data: success")
	return mapResponse, nil
}
func (u UserServiceImpl) GetUserByEmail(email string) (*models.User, error) {
	return u.UserRepo.GetUserByEmail(email)
}

func (u UserServiceImpl) UploadImage(id int, path string) (*models.User, string, error) {
	valUser, err := u.UserRepo.FindById(id)
	if valUser == nil {
		return nil, "Id user not found", nil
	}
	user, err := u.UserRepo.UploadImage(id, path)
	return user, "", err
}
