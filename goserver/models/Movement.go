package models

import (
	"fmt"
	"time"

	"github.com/AcuVuz/barriers-server/classes"
	"github.com/jmoiron/sqlx"
)

type MovementModel struct {
	DB *sqlx.DB
}

func (m MovementModel) InsertForStudent(trx *sqlx.Tx, idBuilding int64, idEvent int64, idStudent int64) (int64, error) {
	var id int64 = 0

	err := trx.Get(&id, `INSERT INTO barriers.moves (id_building, id_event, id_student, event_time) 
		VALUES ($1, $2, $3, $4) RETURNING id`,
		idBuilding, idEvent, idStudent, time.Now())
	if err != nil {
		err := trx.Rollback()
		if err != nil {
			return 0, err
		}
		return 0, err
	}

	return id, nil
}

func (m MovementModel) InsertForEmployee(trx *sqlx.Tx, idBuilding int64, idEvent int64, idEmployee int64) (int64, error) {
	var id int64

	err := trx.Get(
		&id,
		`INSERT INTO barriers.moves (id_building, id_event, id_employee, event_time) 
		VALUES($1, $2, $3, $4) RETURNING id`,
		idBuilding, idEvent, idEmployee, time.Now())
	if err != nil {
		err = trx.Rollback()
		if err != nil {
			return 0, err
		}
		return 0, err
	}

	return id, nil
}

func (m MovementModel) GetMovements(from time.Time, to time.Time) ([]classes.Movement, error) {
	var movements []classes.Movement

	err := m.DB.Select(
		&movements,
		`SELECT m.id, m.id_building, m.id_event, e.name as event_name, m.event_time, 
		m.id_student, m.id_employee
		FROM barriers.moves m
		JOIN barriers.events AS e ON e.id = m.id_event 
		WHERE event_time >= $1 AND event_time <= $2
		ORDER BY event_time DESC`,
		from, to)
	if err != nil {
		fmt.Println(err)
		return movements, err
	}
	return movements, nil
}

func (m MovementModel) GetMovementsForEmployee(idEmployee int64, from time.Time, to time.Time) ([]classes.EmployeeMovement, error) {
	var movements []classes.EmployeeMovement

	err := m.DB.Select(
		&movements,
		`SELECT m.id, m.id_building, m.id_event, e.name as event_name, m.event_time, 
		m.id_employee, p.firstname as lastname, p.name as firstname, p.lastname as middlename, p.skud_card
		FROM barriers.moves m
		JOIN barriers.events AS e ON  e.id = m.id_event
		LEFT JOIN pers."Persons" as p ON m.id_employee = p.id
		WHERE id_employee = $1 AND (event_time >= $2 AND event_time <= $3)`,
		idEmployee, from, to)
	if err != nil {
		fmt.Println(err)
		return movements, err
	}
	return movements, nil
}

func (m MovementModel) GetMovementsForStudent(idStudent int64, from time.Time, to time.Time) ([]classes.StudentMovement, error) {
	var movements []classes.StudentMovement

	err := m.DB.Select(
		&movements,
		`SELECT m.id, m.id_building, m.id_event, e.name as event_name, m.event_time, 
		m.id_student, s.firstname, s.middlename, s.lastname, s.skud_card  
		FROM barriers.moves m
		JOIN barriers.events AS e ON e.id = m.id_event
		JOIN education.students AS s ON s.id = m.id_student
		WHERE id_student = $1 AND (event_time >= $2 AND event_time <= $3)`,
		idStudent, from, to)
	if err != nil {
		fmt.Println(err)
		return movements, err
	}
	return movements, nil
}
