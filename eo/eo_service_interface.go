package eo

import "sieo_app/models"

type EoServiceInterface interface {
	ReadAll() (map[string]interface{}, error)
	Insert(req *models.Eo) (map[string]interface{}, error)
	UpdateEo(id int, req *models.Eo, user *models.User) (map[string]interface{}, map[string]interface{}, error)
}
