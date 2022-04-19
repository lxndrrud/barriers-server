package models

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type Student struct {
	Id         int64  `db:"id"`
	Firstname  string `db:"firstname"`
	Middlename string `db:"middlename"`
	Lastname   string `db:"lastname"`
	SkudCard   string `db:"skud_card"`
}

type Document struct {
	TypeDocumentName string
	Series           string
	Number           string
	Issued           string
	IssuedDate       string
}

type StudentModel struct {
	DB *sqlx.DB
}

func (m StudentModel) Get(id int64) (Student, error) {
	var student Student
	err := m.DB.Get(
		&student,
		`SELECT id, firstname, middlename, lastname, skud_card FROM "education"."students" 
			WHERE id = $1`,
		id)
	if err == sql.ErrNoRows {
		return Student{}, nil
	}
	if err != nil {
		return Student{}, err
	}
	return student, nil
}

func (m StudentModel) GetBySkudCard(SkudCard string) (Student, error) {
	var student Student
	err := m.DB.Get(
		&student,
		`SELECT id, firstname, middlename, lastname, skud_card FROM "education"."students" 
			WHERE skud_card = $1`,
		SkudCard)
	if err == sql.ErrNoRows {
		return Student{}, nil
	}
	if err != nil {
		return Student{}, err
	}
	return student, nil
}
