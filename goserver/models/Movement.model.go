package models

import (
	"fmt"
	"time"

	"github.com/AcuVuz/barriers-server/classes"
	"github.com/AcuVuz/barriers-server/utils"
	"github.com/jmoiron/sqlx"
)

type MovementModel struct {
	DB *sqlx.DB
}

func (m MovementModel) Insert(trx *sqlx.Tx, idBuilding, idEvent,
	idEmployee, idStudent int64) (int64, error) {
	var id int64 = 0
	err := trx.Get(&id, `INSERT INTO barriers.moves (id_building, id_event, id_student, id_employee, event_time) 
			VALUES ($1, $2, $3, $4, $5) RETURNING id`,
		idBuilding, idEvent, utils.NewNullInt(idStudent), utils.NewNullInt(idEmployee), time.Now())
	if err != nil {
		err := trx.Rollback()
		if err != nil {
			return 0, err
		}
		return 0, err
	}

	return id, nil
}

func (m MovementModel) GetMovements(from time.Time, to time.Time) ([]classes.DatabaseMovement, error) {
	movements := make([]classes.DatabaseMovement, 0)

	err := m.DB.Select(
		&movements,
		`SELECT m.id, m.id_building, m.id_event, e.name as event_name, m.event_time, 
			m.id_student, m.id_employee, b.name as building_name,
			p.firstname as employee_firstname, p.name as employee_middlename, 
			p.lastname as employee_lastname, p.skud_card as employee_skud_card,
			s.firstname as student_firstname, s.middlename as student_middlename, 
			s.lastname as student_lastname, s.skud_card as student_skud_card
		FROM barriers.moves m
		JOIN barriers.events AS e ON e.id = m.id_event 
		JOIN barriers.buildings AS b ON b.id = m.id_building
		LEFT JOIN pers."Persons" AS p ON p.id = m.id_employee
		LEFT JOIN education.students AS s ON s.id = m.id_student
		WHERE m.event_time >= $1 AND m.event_time <= $2
		ORDER BY event_time DESC`,
		from, to)
	if err != nil {
		fmt.Println(err)
		return movements, err
	}
	return movements, nil
}

func (m MovementModel) GetMovementsForBuilding(idBuilding int64, from time.Time,
	to time.Time) ([]classes.DatabaseMovement, error) {

	movements := make([]classes.DatabaseMovement, 0)

	err := m.DB.Select(
		&movements,
		`SELECT m.id, m.id_building, m.id_event, e.name as event_name, m.event_time, 
			m.id_student, m.id_employee, b.name as building_name,
			p.firstname as employee_firstname, p.name as employee_middlename, 
			p.lastname as employee_lastname, p.skud_card as employee_skud_card,
			s.firstname as student_firstname, s.middlename as student_middlename, 
			s.lastname as student_lastname, s.skud_card as student_skud_card
		FROM barriers.moves m
		JOIN barriers.events AS e ON e.id = m.id_event 
		JOIN barriers.buildings AS b ON b.id = m.id_building
		LEFT JOIN pers."Persons" AS p ON p.id = m.id_employee
		LEFT JOIN education.students AS s ON s.id = m.id_student
		WHERE m.id_building = $1 AND (m.event_time >= $2 AND m.event_time <= $3)
		ORDER BY event_time DESC`,
		idBuilding, from, to)
	if err != nil {
		fmt.Println(err)
		return movements, err
	}
	return movements, nil
}

func (m MovementModel) GetMovementsForUser(idStudent, idEmployee int64, from time.Time, to time.Time) ([]classes.Movement, error) {
	movements := make([]classes.Movement, 0)

	if idStudent != 0 {
		err := m.DB.Select(
			&movements,
			`SELECT m.id, m.id_building, m.id_event, m.event_time, m.id_student, m.id_employee,
			e.name as event_name,
			b.name as building_name
			FROM barriers.moves m
			JOIN barriers.events AS e ON e.id = m.id_event
			JOIN barriers.buildings AS b ON b.id = m.id_building
			WHERE m.id_student = $1 AND (event_time >= $2 AND event_time <= $3)
			ORDER BY event_time DESC`,
			idStudent, from, to)
		if err != nil {
			fmt.Println(err)
			return movements, err
		}

	} else if idEmployee != 0 {
		err := m.DB.Select(
			&movements,
			`SELECT m.id, m.id_building, m.id_event, m.event_time, m.id_student, m.id_employee,
			e.name as event_name,
			b.name as building_name
			FROM barriers.moves m
			JOIN barriers.events AS e ON e.id = m.id_event
			JOIN barriers.buildings AS b ON b.id = m.id_building
			WHERE m.id_employee = $1 AND (event_time >= $2 AND event_time <= $3)
			ORDER BY event_time DESC`,
			idEmployee, from, to)
		if err != nil {
			fmt.Println(err)
			return movements, err
		}
	} else {
		// ???????????????????? ??????????
		err := m.DB.Select(
			&movements,
			`SELECT m.id, m.id_building, m.id_event, m.event_time, m.id_student, m.id_employee,
			e.name as event_name,
			b.name as building_name
			FROM barriers.moves m
			JOIN barriers.events AS e ON e.id = m.id_event
			JOIN barriers.buildings AS b ON b.id = m.id_building
			WHERE m.id_employee is NULL AND m.id_student is NULL AND (event_time >= $1 AND event_time <= $2)
			ORDER BY event_time DESC`,
			from, to)
		if err != nil {
			fmt.Println(err)
			return movements, err
		}
	}

	return movements, nil

}

func (m MovementModel) GetMovementsForUserBuilding(idBuilding, idStudent, idEmployee int64, from time.Time, to time.Time) ([]classes.Movement, error) {
	movements, err := m.GetMovementsForUser(idStudent, idEmployee, from, to)
	if err != nil {
		return movements, err
	}

	movementsResult := make([]classes.Movement, 0)
	for _, m2 := range movements {
		fmt.Println(m2.IdBuilding == idBuilding)
		if m2.IdBuilding == idBuilding {
			movementsResult = append(movementsResult, m2)
		}
	}
	fmt.Println(movementsResult)

	return movementsResult, nil
}
