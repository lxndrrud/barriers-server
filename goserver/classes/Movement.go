package classes

import (
	"database/sql"
	"time"
)

type Movement struct {
	Id             int64         `db:"id" json:"id"`
	IdBuilding     int64         `db:"id_building" json:"id_building"`
	IdEvent        int64         `db:"id_event" json:"id_event"`
	EventName      string        `db:"event_name" json:"event_name"`
	EventTimestamp time.Time     `db:"event_time" json:"event_timestamp"`
	IdStudent      sql.NullInt64 `db:"id_student" json:"id_student"`
	IdEmployee     sql.NullInt64 `db:"id_employee" json:"id_employee"`
}

type MovementJSON struct {
	Id             int64     `json:"id"`
	IdBuilding     int64     `json:"id_building"`
	IdEvent        int64     `json:"id_event"`
	EventName      string    `json:"event_name"`
	EventTimestamp time.Time `json:"event_timestamp"`
	IdStudent      int64     `json:"id_student"`
	IdEmployee     int64     `json:"id_employee"`
}

type DatabaseStudentMovement struct {
	Movement
	UserBase
}

type JSONStudentMovement struct {
	MovementJSON
	UserJSONBase
}

func CreateJSONFromStudentMovement(dbMovement *DatabaseStudentMovement) JSONStudentMovement {
	return JSONStudentMovement{
		UserJSONBase: UserJSONBase{
			Firstname:  dbMovement.Firstname,
			Middlename: dbMovement.Middlename,
			Lastname:   dbMovement.Lastname,
			SkudCard:   dbMovement.SkudCard.String,
		},
		MovementJSON: MovementJSON{
			Id:             dbMovement.Id,
			IdEvent:        dbMovement.IdEvent,
			IdBuilding:     dbMovement.IdBuilding,
			IdStudent:      0,
			IdEmployee:     dbMovement.IdEmployee.Int64,
			EventName:      dbMovement.EventName,
			EventTimestamp: dbMovement.EventTimestamp,
		},
	}
}

type DatabaseEmployeeMovement struct {
	Movement
	UserBase
}

type JSONEmployeeMovement struct {
	MovementJSON
	UserJSONBase
}

func CreateJSONFromEmployeeMovement(dbMovement *DatabaseEmployeeMovement) JSONEmployeeMovement {
	return JSONEmployeeMovement{
		UserJSONBase: UserJSONBase{
			Firstname:  dbMovement.Firstname,
			Middlename: dbMovement.Middlename,
			Lastname:   dbMovement.Lastname,
			SkudCard:   dbMovement.SkudCard.String,
		},
		MovementJSON: MovementJSON{
			Id:             dbMovement.Id,
			IdEvent:        dbMovement.IdEvent,
			IdBuilding:     dbMovement.IdBuilding,
			IdStudent:      0,
			IdEmployee:     dbMovement.IdEmployee.Int64,
			EventName:      dbMovement.EventName,
			EventTimestamp: dbMovement.EventTimestamp,
		},
	}
}
