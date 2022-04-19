package models

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

type Movement struct {
	Id             int64         `db:"id"`
	IdBuilding     int64         `db:"id_building"`
	IdEvent        int64         `db:"id_event"`
	EventName      string        `db:"event_name"`
	EventTimestamp time.Time     `db:"event_time"`
	IdStudent      sql.NullInt64 `db:"id_student"`
	IdEmployee     sql.NullInt64 `db:"id_employee"`
}

type StudentMovement struct {
	Id             int64     `db:"id"`
	IdBuilding     int64     `db:"id_building"`
	IdEvent        int64     `db:"id_event"`
	EventName      string    `db:"event_name"`
	EventTimestamp time.Time `db:"event_time"`
	IdStudent      int64     `db:"id_student"`
}

type EmployeeMovement struct {
	Id             int64     `db:"id"`
	IdBuilding     int64     `db:"id_building"`
	IdEvent        int64     `db:"id_event"`
	EventName      string    `db:"event_name"`
	EventTimestamp time.Time `db:"event_time"`
	IdEmployee     int64     `db:"id_employee"`
}

type MovementModel struct {
	DB *sqlx.DB
}

func (m MovementModel) InsertForStudent(trx *sqlx.Tx, idBuilding int64, idEvent int64, idStudent int64) (int64, error) {
	var id int64 = 0

	err := trx.Get(&id, `INSERT INTO "barriers"."moves" (id_building, id_event, id_student, event_time) 
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
		`INSERT INTO "barriers"."moves" (id_building, id_event, id_employee, event_time) 
		VALUES($1, $2, $3, $4) RETURNING id`,
		idBuilding, idEvent, idEmployee, time.Now())
	if err != nil {
		fmt.Println("rolledback", err)
		err = trx.Rollback()
		if err != nil {
			return 0, err
		}
		return 0, err
	}

	return id, nil
}

func (m MovementModel) GetMovements(from time.Time, to time.Time) ([]Movement, error) {
	var movements []Movement

	err := m.DB.Select(
		&movements,
		`SELECT moves.id, id_building, id_event, e.name as event_name, event_time, id_student, id_employee  FROM "barriers"."moves" 
		JOIN barriers.events as e ON e.id = barriers.moves.id_event 
		WHERE event_time >= $1 AND event_time <= $2`,
		from, to)
	if err != nil {
		fmt.Println(err)
		return movements, err
	}
	return movements, nil
}

func (m MovementModel) GetMovementsForEmployee(idEmployee int64, from time.Time, to time.Time) ([]EmployeeMovement, error) {
	var movements []EmployeeMovement

	err := m.DB.Select(
		&movements,
		`SELECT id, id_building, id_event, event_time, id_employee FROM "barriers"."moves" 
		WHERE id_employee = $1 AND (event_time >= $2 AND event_time <= $3)`,
		idEmployee, from, to)
	if err != nil {
		return movements, err
	}
	return movements, nil
}

func (m MovementModel) GetMovementsForStudent(idStudent int64, from time.Time, to time.Time) ([]StudentMovement, error) {
	var movements []StudentMovement

	err := m.DB.Select(
		&movements,
		`SELECT id, id_building, id_event, event_time, id_student  FROM "barriers"."moves" 
		WHERE id_student = $1 AND (event_time >= $2 AND event_time <= $3)`,
		idStudent, from, to)
	if err != nil {
		return movements, err
	}
	return movements, nil
}
