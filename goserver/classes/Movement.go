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

type StudentMovement struct {
	Movement
	UserBase
}

type EmployeeMovement struct {
	Movement
	UserBase
}
