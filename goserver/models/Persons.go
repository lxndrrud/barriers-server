package models

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type Person struct {
	Id         int    `db:"id"`
	Firstname  string `db:"firstname"`
	Middlename string `db:"name"`
	Lastname   string `db:"lastname"`
	SkudCard   string `db:"skud_card"`
}

type PersonModel struct {
	DB *sqlx.DB
}

func (m PersonModel) GetBySkudCard(SkudCard string) (Person, error) {
	var person Person

	err := m.DB.Get(&person,
		`SELECT id, firstname, name, lastname, skud_card FROM "pers"."Persons"
			WHERE skud_card = $1`,
		SkudCard)
	if err == sql.ErrNoRows {
		return Person{}, nil
	}
	if err != nil {
		return Person{}, err
	}
	return person, nil
}
