package service

import (
	"fmt"
	"sieo_app/common"
	"sieo_app/eo"
	"sieo_app/models"
	"sieo_app/user"

	"github.com/sirupsen/logrus"
)

type EoServiceImpl struct {
	EoRepository eo.EoRepoInterface
	UserRepo     user.RepoInterface
}

func CreateEoService(EoRepository eo.EoRepoInterface, UserRepo user.RepoInterface) eo.EoServiceInterface {
	return &EoServiceImpl{EoRepository: EoRepository, UserRepo: UserRepo}
}

func (e *EoServiceImpl) ReadAll() (map[string]interface{}, error) {
	response, err := e.EoRepository.ReadAll()
	if err != nil {
		logrus.Error(err)
		return common.Message(false, "Opps Something Wrong"), err
	}

	mapResponse := common.Message(true, "Read All EO Data: Success")
	mapResponse["response"] = response
	return mapResponse, nil
}

func (e *EoServiceImpl) Insert(req *models.Eo) (map[string]interface{}, error) {
	response, err := e.EoRepository.Insert(req)
	if err != nil {
		logrus.Error(err)
		return common.Message(false, "Opps Something Wrong"), err
	}

	mapResponse := common.Message(true, "Insert Eo Data : Successs")
	mapResponse["response"] = response
	return mapResponse, nil
}

func (e *EoServiceImpl) UpdateEo(id int, req *models.Eo, user *models.User) (map[string]interface{}, error) {
	response, err := e.EoRepository.UpdateEo(id, req)
	if err != nil {
		logrus.Error(err)
		return common.Message(false, "Opps Something Wrong"), err
	}
	_, err = e.UserRepo.Update(id, user)
	if err != nil {
		logrus.Error(err)
		return common.Message(false, "Opps Something Wrong User"), err
	}
	fmt.Println("Keprint")
	mapResponse := common.Message(true, "Update Eo Data  : Success")
	mapResponse["response"] = response
	return mapResponse, nil
}
