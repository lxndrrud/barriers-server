package models

import (
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"
)

type Movement struct {
	Id             int           `db:"id"`
	IdBuilding     int           `db:"id_building"`
	IdEvent        int           `db:"id_event"`
	EventTimestamp time.Time     `db:"event_time"`
	IdStudent      sql.NullInt64 `db:"id_student"`
	IdEmployee     sql.NullInt64 `db:"id_employee"`
}

type MovementModel struct {
	DB *sqlx.DB
}

func (m MovementModel) InsertForStudent(idBuilding int64, idEvent int64, idStudent int64) (int64, error) {
	var id int64 = 0

	trx, err := m.DB.Beginx()
	if err != nil {
		return 0, nil
	}

	err = trx.Get(&id, `INSERT INTO "barriers"."moves" (id_building, id_event, id_student, event_time) 
		VALUES ($1, $2, $3, $4) RETURNING id`,
		idBuilding, idEvent, idStudent, time.Now())
	if err != nil {
		err := trx.Rollback()
		if err != nil {
			return 0, err
		}
		return 0, err
	}

	err = trx.Commit()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (m MovementModel) InsertForEmployee(idBuilding int64, idEvent int64, idEmployee int64) (int64, error) {
	var id int64

	trx, err := m.DB.Beginx()
	if err != nil {
		return 0, nil
	}

	err = trx.Get(&id, `INSERT INTO 'barriers'.'moves' (id_building, id_event, id_employee, event_time) 
		VALUES($1, $2, $3, $4) RETURNING 'barriers'.'moves'.'id'`,
		idBuilding, idEvent, idEmployee, time.Now())
	if err != nil {
		err = trx.Rollback()
		if err != nil {
			return 0, err
		}
		return 0, err
	}

	err = trx.Commit()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (m MovementModel) GetMovements(from time.Time, to time.Time) ([]Movement, error) {
	var movements []Movement

	err := m.DB.Select(&movements, `SELECT * FROM "barriers"."moves" 
		WHERE event_time >= $1 AND event_time <= to`,
		from, to)
	if err != nil {
		return movements, err
	}
	return movements, nil
}
