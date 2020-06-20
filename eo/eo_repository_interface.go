package eo

import "sieo_app/models"

type EoRepoInterface interface {
	ReadAll() ([]*models.Eo, error)
	Insert(req *models.Eo) (*models.Eo, error)
	UpdateEo(id int, req *models.Eo) (*models.Eo, error)
}
