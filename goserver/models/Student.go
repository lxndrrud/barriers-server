package models

import (
	"database/sql"

	"github.com/AcuVuz/barriers-server/classes"
	"github.com/jmoiron/sqlx"
)

type StudentModel struct {
	DB *sqlx.DB
}

func (m StudentModel) Get(id int64) (classes.Student, error) {
	var student classes.Student
	err := m.DB.Get(
		&student,
		`SELECT id, firstname, middlename, lastname, skud_card FROM education.students 
			WHERE id = $1`,
		id)
	if err == sql.ErrNoRows {
		return classes.Student{}, nil
	}
	if err != nil {
		return classes.Student{}, err
	}
	return student, nil
}

func (m StudentModel) GetBySkudCard(SkudCard string) (classes.Student, error) {
	var student classes.Student
	err := m.DB.Get(
		&student,
		`SELECT id, firstname, middlename, lastname, skud_card FROM education.students 
			WHERE skud_card = $1`,
		SkudCard)
	if err == sql.ErrNoRows {
		return classes.Student{}, nil
	}
	if err != nil {
		return classes.Student{}, err
	}
	return student, nil
}
