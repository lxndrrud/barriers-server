package interfaces

type UserBase struct {
	Id         int    `db:"id"`
	Firstname  string `db:"firstname"`
	Middlename string `db:"middlename"`
	Lastname   string `db:"lastname"`
}
