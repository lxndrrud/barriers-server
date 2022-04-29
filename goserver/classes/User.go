package classes

import "database/sql"

type UserBase struct {
	Firstname  string         `db:"firstname"`
	Middlename string         `db:"middlename"`
	Lastname   string         `db:"lastname"`
	SkudCard   sql.NullString `db:"skud_card"`
}

type UserJSONBase struct {
	Firstname  string `json:"firstname"`
	Middlename string `json:"middlename"`
	Lastname   string `json:"lastname"`
	SkudCard   string `json:"skud_card"`
}

type User struct {
	Id int64 `db:"id"`
	UserBase
}

type Student struct {
	User
}

type Employee struct {
	User
}

type UserJSON struct {
	Id   int64  `json:"id"`
	Type string `json:"type"`
	UserJSONBase
}

func CreateUserJSONFromStudent(student *Student) UserJSON {
	return UserJSON{
		Id:   student.Id,
		Type: "student",
		UserJSONBase: UserJSONBase{
			Firstname:  student.Firstname,
			Lastname:   student.Lastname,
			Middlename: student.Middlename,
			SkudCard:   student.SkudCard.String,
		},
	}
}

func CreateUserJSONFromEmployee(employee *Employee) UserJSON {
	return UserJSON{
		Id:   employee.Id,
		Type: "employee",
		UserJSONBase: UserJSONBase{
			Firstname:  employee.Firstname,
			Lastname:   employee.Lastname,
			Middlename: employee.Middlename,
			SkudCard:   employee.SkudCard.String,
		},
	}
}
