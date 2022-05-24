package classes

import "database/sql"

type DBStudentPersonalInfo struct {
	Id         int64          `db:"id"`
	Firstname  string         `db:"firstname"`
	Middlename string         `db:"middlename"`
	Lastname   string         `db:"lastname"`
	SkudCard   sql.NullString `db:"skud_card"`
}

type JSONStudentPersonalInfo struct {
	Id         int64  `json:"id"`
	Firstname  string `json:"firstname"`
	Middlename string `json:"middlename"`
	Lastname   string `json:"lastname"`
	SkudCard   string `json:"skud_card"`
}

func CreateJSONStudentPersonalInfo(s *DBStudentPersonalInfo) JSONStudentPersonalInfo {
	newValue := JSONStudentPersonalInfo{
		Id:         s.Id,
		Firstname:  s.Firstname,
		Middlename: s.Middlename,
		Lastname:   s.Lastname,
	}
	if s.SkudCard.Valid {
		newValue.SkudCard = s.SkudCard.String
	}
	return newValue
}

type DBStudentGroupInfo struct {
	Id              int64  `db:"id"`
	Title           string `db:"title"`
	Course          string `db:"course"`
	DepartmentTitle string `db:"department_title"`
}

type JSONStudentGroupInfo struct {
	Id              int64  `json:"id"`
	Title           string `json:"title"`
	Course          string `json:"course"`
	DepartmentTitle string `json:"department_title"`
}

func CreateJSONStudentGroupInfo(g []DBStudentGroupInfo) []JSONStudentGroupInfo {
	result := make([]JSONStudentGroupInfo, 0)

	for _, value := range g {
		result = append(result, JSONStudentGroupInfo{
			Id:              value.Id,
			Title:           value.Title,
			Course:          value.Course,
			DepartmentTitle: value.DepartmentTitle,
		})
	}

	return result
}
