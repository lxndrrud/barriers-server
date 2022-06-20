package models

import (
	"github.com/AcuVuz/barriers-server/classes"
	"github.com/jmoiron/sqlx"
)

type BuildingModel struct {
	DB *sqlx.DB
}

func (m BuildingModel) GetAll() ([]classes.Building, error) {
	buildings := make([]classes.Building, 0)
	err := m.DB.Select(&buildings, `SELECT id, name FROM barriers.buildings`)
	if err != nil {
		return buildings, err
	}
	buildings = append(buildings, classes.Building{
		Id:   0,
		Name: "Все корпуса",
	})
	return buildings, nil
}
