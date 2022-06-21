package classes

import (
	"database/sql"
	"encoding/json"
	"strconv"
)

type DBUser struct {
	Id         int64          `db:"id"`
	Firstname  string         `db:"firstname"`
	Middlename string         `db:"middlename"`
	Lastname   string         `db:"lastname"`
	DBSkudCard sql.NullString `db:"skud_card"`
	Type       string         `json:"type"`
}

func (c DBUser) MarshalJSON() ([]byte, error) {
	var skudCard string
	if c.DBSkudCard.Valid {
		skudCard = c.DBSkudCard.String
	} else {
		skudCard = ""
	}
	m := map[string]string{
		"id":         strconv.FormatInt(c.Id, 10),
		"firstname":  c.Firstname,
		"middlename": c.Middlename,
		"lastname":   c.Lastname,
		"skud_card":  skudCard,
		"type":       c.Type,
	}

	return json.Marshal(m)
}

type DBStudentGroupInfo struct {
	Id              int64  `db:"id" json:"id"`
	Title           string `db:"title" json:"title"`
	Course          string `db:"course" json:"course"`
	DepartmentTitle string `db:"department_title" json:"department_title"`
}

type DBEmployeePositionInfo struct {
	Id              int64          `db:"id"`
	Title           string         `db:"title"`
	DepartmentTitle string         `db:"department_title"`
	DBDateDrop      sql.NullString `db:"date_drop"`
}

func (c DBEmployeePositionInfo) MarshalJSON() ([]byte, error) {
	var dateDrop string
	if c.DBDateDrop.Valid {
		dateDrop = c.DBDateDrop.String
	} else {
		dateDrop = "Все еще работает"
	}
	m := map[string]string{
		"id":               strconv.FormatInt(c.Id, 10),
		"title":            c.Title,
		"department_title": c.DepartmentTitle,
		"date_drop":        dateDrop,
	}
	return json.Marshal(m)
}

type JSONEmployee struct {
	Employee  DBUser                   `json:"employee"`
	Positions []DBEmployeePositionInfo `json:"positions"`
}

type JSONStudent struct {
	Student DBUser               `json:"student"`
	Groups  []DBStudentGroupInfo `json:"groups"`
}
