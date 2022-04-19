package models

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type Employee struct {
	Id         int64  `db:"id"`
	Firstname  string `db:"firstname"`
	Middlename string `db:"name"`
	Lastname   string `db:"lastname"`
	SkudCard   string `db:"skud_card"`
}

type EmployeeModel struct {
	DB *sqlx.DB
}

func (m EmployeeModel) Get(id int64) (Employee, error) {
	var person Employee

	err := m.DB.Get(
		&person,
		`SELECT id, firstname, name, lastname, skud_card FROM "pers"."Persons"
			WHERE id = $1`,
		id)
	if err == sql.ErrNoRows {
		return Employee{}, nil
	}
	if err != nil {
		return Employee{}, err
	}
	return person, nil
}

func (m EmployeeModel) GetBySkudCard(SkudCard string) (Employee, error) {
	var person Employee

	err := m.DB.Get(
		&person,
		`SELECT id, firstname, name, lastname, skud_card FROM "pers"."Persons"
			WHERE skud_card = $1`,
		SkudCard)
	if err == sql.ErrNoRows {
		return Employee{}, nil
	}
	if err != nil {
		return Employee{}, err
	}
	return person, nil
}
