package classes

import "database/sql"

type UserBase struct {
	Firstname  string         `db:"firstname" json:"firstname"`
	Middlename string         `db:"middlename" json:"middlename"`
	Lastname   string         `db:"lastname" json:"lastname"`
	SkudCard   sql.NullString `db:"skud_card" json:"skud_card"`
}

type User struct {
	Id int64 `db:"id" json:"id"`
	UserBase
}

type Student struct {
	User
}

type Employee struct {
	User
}
