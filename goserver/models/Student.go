package models

import (
	"database/sql"

	"github.com/AcuVuz/barriers-server/interfaces"
	"github.com/jmoiron/sqlx"
)

type Student struct {
	interfaces.UserBase
	SkudCard string `db:"skud_card"`
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

func (m StudentModel) GetBySkudCard(SkudCard string) (Student, error) {
	var student Student
	err := m.DB.Get(&student,
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
