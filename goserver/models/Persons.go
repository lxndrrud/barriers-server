package models

import (
	"database/sql"

	"github.com/AcuVuz/barriers-server/interfaces"
	"github.com/jmoiron/sqlx"
)

type Person struct {
	interfaces.UserBase
	SkudCard string `db:"skud_card"`
}

type PersonModel struct {
	DB *sqlx.DB
}

func (m PersonModel) GetBySkudCard(SkudCard string) (Person, error) {
	var person Person

	err := m.DB.Get(&person,
		`SELECT id, firstname, middlename, lastname, skud_card FROM "pers"."Persons"
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
