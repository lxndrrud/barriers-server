package models

import (
	"database/sql"

	"github.com/AcuVuz/barriers-server/classes"
	"github.com/jmoiron/sqlx"
)

type EmployeeModel struct {
	DB *sqlx.DB
}

func (m EmployeeModel) Get(id int64) (classes.Employee, error) {
	var person classes.Employee

	err := m.DB.Get(
		&person,
		`SELECT id, firstname as lastname, name as firstname, lastname as middlename, skud_card FROM pers."Persons"
			WHERE id = $1`,
		id)
	if err == sql.ErrNoRows {
		return classes.Employee{}, nil
	}
	if err != nil {
		return classes.Employee{}, err
	}
	return person, nil
}

func (m EmployeeModel) GetBySkudCard(SkudCard string) (classes.Employee, error) {
	var person classes.Employee

	err := m.DB.Get(
		&person,
		`SELECT id, firstname as lastname, name as firstname, lastname as middlename, skud_card FROM pers."Persons"
			WHERE skud_card = $1`,
		SkudCard)
	if err == sql.ErrNoRows {
		return classes.Employee{}, nil
	}
	if err != nil {
		return classes.Employee{}, err
	}
	return person, nil
}
