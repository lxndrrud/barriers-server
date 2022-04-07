package models

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type Student struct {
	Id         int            `db:"id"`
	Firstname  string         `db:"firstname"`
	Middlename string         `db:"middlename"`
	Lastname   string         `db:"lastname"`
	SkudCard   sql.NullString `db:"skud_card"`
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

func (m StudentModel) GetBySkudCard(SkudCard string) (*Student, error) {
	student := Student{}
	err := m.DB.Get(&student,
		`SELECT id, firstname, middlename, lastname FROM "education"."students" 
			WHERE skud_card = $1`,
		SkudCard)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &student, nil
}
