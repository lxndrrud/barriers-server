package classes

import "database/sql"

type DBEmployeePersonalInfo struct {
	Id         int64          `db:"id"`
	Firstname  string         `db:"firstname"`
	Middlename string         `db:"middlename"`
	Lastname   string         `db:"lastname"`
	SkudCard   sql.NullString `db:"skud_card"`
}

type JSONEmployeePersonalInfo struct {
	Id         int64  `json:"id"`
	Firstname  string `json:"firstname"`
	Middlename string `json:"middlename"`
	Lastname   string `json:"lastname"`
	SkudCard   string `json:"skud_card"`
}

func CreateJSONEmployeePersonalInfo(c *DBEmployeePersonalInfo) JSONEmployeePersonalInfo {
	return JSONEmployeePersonalInfo{
		Id:         c.Id,
		Firstname:  c.Firstname,
		Middlename: c.Middlename,
		Lastname:   c.Lastname,
		SkudCard:   c.SkudCard.String,
	}
}

type DBEmployeePositionInfo struct {
	Id              int64          `db:"id"`
	Title           string         `db:"title"`
	DepartmentTitle string         `db:"department_title"`
	DateDrop        sql.NullString `db:"date_drop"`
}

type JSONEmployeePositionInfo struct {
	Id              int64  `json:"id"`
	Title           string `json:"title"`
	DepartmentTitle string `json:"department_title"`
	DateDrop        string `json:"date_drop"`
}

func CreateJSONEmployeePositionInfo(c []DBEmployeePositionInfo) []JSONEmployeePositionInfo {
	result := make([]JSONEmployeePositionInfo, 0)

	for _, value := range c {
		newValue := JSONEmployeePositionInfo{
			Id:              value.Id,
			Title:           value.Title,
			DepartmentTitle: value.DepartmentTitle,
		}
		if value.DateDrop.Valid {
			newValue.DateDrop = value.DateDrop.String
		} else {
			newValue.DateDrop = "Всё еще работает"
		}
		result = append(result, newValue)
	}

	return result
}
