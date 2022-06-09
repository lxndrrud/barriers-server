package models

import (
	"database/sql"

	"github.com/AcuVuz/barriers-server/classes"
	"github.com/jmoiron/sqlx"
)

type EmployeeModel struct {
	DB *sqlx.DB
}

func (m EmployeeModel) Get(id int64) (classes.DBUser, error) {
	var employee classes.DBUser

	err := m.DB.Get(
		&employee,
		`SELECT id, firstname as lastname, name as firstname, lastname as middlename, skud_card FROM pers."Persons"
			WHERE id = $1`,
		id)

	if err != nil {
		return classes.DBUser{}, err
	}
	employee.Type = "Сотрудник"
	return employee, nil
}

func (m EmployeeModel) GetBySkudCard(SkudCard string) (classes.DBUser, error) {
	var employee classes.DBUser

	err := m.DB.Get(
		&employee,
		`SELECT id, firstname as lastname, name as firstname, lastname as middlename, skud_card FROM pers."Persons"
			WHERE skud_card = $1`,
		SkudCard)
	if err == sql.ErrNoRows {
		return classes.DBUser{}, nil
	}
	if err != nil {
		return classes.DBUser{}, err
	}
	employee.Type = "Сотрудник"
	return employee, nil
}

func (m EmployeeModel) GetPositionsInfo(IdEmployee int64) ([]classes.DBEmployeePositionInfo1, error) {
	positionsInfo := make([]classes.DBEmployeePositionInfo1, 0)

	err := m.DB.Select(
		&positionsInfo,
		`SELECT pos.id, typepos.name_position as title, name_department as department_title, 
			date_drop FROM pers."PersonsPosition" perspos
		JOIN pers."Position" AS pos ON pos.id = perspos.id_position
		JOIN pers."Departments" AS dep ON dep.id = pos.id_department
		JOIN pers."TypePositions" AS typepos ON typepos.id = pos.id_type
		WHERE perspos.id_person = $1`,
		IdEmployee)

	if err != nil {
		return positionsInfo, err
	}

	return positionsInfo, nil
}
