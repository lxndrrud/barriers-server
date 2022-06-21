package models

import (
	"database/sql"

	"github.com/AcuVuz/barriers-server/classes"
	"github.com/jmoiron/sqlx"
)

type StudentModel struct {
	DB *sqlx.DB
}

func (m StudentModel) Get(id int64) (classes.DBUser, error) {
	var student classes.DBUser
	err := m.DB.Get(
		&student,
		`SELECT id, firstname, middlename, lastname, skud_card FROM education.students 
			WHERE id = $1`,
		id)
	if err != nil {
		return student, err
	}
	student.Type = "Студент"
	return student, nil
}

func (m StudentModel) GetBySkudCard(SkudCard string) (classes.DBUser, error) {
	var student classes.DBUser
	err := m.DB.Get(
		&student,
		`SELECT id, firstname, middlename, lastname, skud_card FROM education.students 
			WHERE skud_card = $1`,
		SkudCard)
	if err == sql.ErrNoRows {
		return student, nil
	}
	if err != nil {
		return student, err
	}
	student.Type = "Студент"
	return student, nil
}

func (m StudentModel) GetGroupsInfo(IdStudent int64) ([]classes.DBStudentGroupInfo, error) {
	groupsList := make([]classes.DBStudentGroupInfo, 0)

	err := m.DB.Select(
		&groupsList,
		`SELECT g.id, g.nickname as title, g.course, 
			dep.name_department as department_title 
		FROM education.students_groups sg
		JOIN education.study_groups AS g ON g.id = sg.id_group 
		LEFT JOIN pers."Departments" AS dep ON dep.id = g.id_faculty
		WHERE sg.id_student = $1`,
		IdStudent)

	if err != nil {
		return groupsList, err
	}
	return groupsList, nil
}
