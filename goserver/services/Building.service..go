package services

import (
	"net/http"

	"github.com/AcuVuz/barriers-server/classes"
	"github.com/AcuVuz/barriers-server/models"
	"github.com/jmoiron/sqlx"
)

type BuildingsService struct {
	db        *sqlx.DB
	buildings interface {
		GetAll() ([]classes.Building, error)
	}
}

func CreateBuildingsService(db *sqlx.DB) *BuildingsService {
	return &BuildingsService{
		db:        db,
		buildings: &models.BuildingModel{DB: db},
	}
}

func (s BuildingsService) GetAll() ([]classes.Building, *classes.CustomError) {
	buildings, err := s.buildings.GetAll()
	if err != nil {
		return buildings, &classes.CustomError{
			Code: http.StatusInternalServerError,
			Text: "Внутренняя ошибка при поиске зданий: " + err.Error(),
		}
	}
	return buildings, nil
}
