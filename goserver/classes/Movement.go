package classes

import (
	"database/sql"
	"encoding/json"
	"time"
)

type Movement struct {
	Id             int64         `db:"id"`
	IdBuilding     int64         `db:"id_building"`
	BuildingName   string        `db:"building_name"`
	IdEvent        int64         `db:"id_event"`
	EventName      string        `db:"event_name"`
	EventTimestamp time.Time     `db:"event_time"`
	IdStudent      sql.NullInt64 `db:"id_student"`
	IdEmployee     sql.NullInt64 `db:"id_employee"`
}

func (c Movement) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"id":              c.Id,
		"id_building":     c.IdBuilding,
		"id_event":        c.IdEvent,
		"building_name":   c.BuildingName,
		"event_name":      c.EventName,
		"event_timestamp": c.EventTimestamp,
		"id_student":      c.IdStudent.Int64,
		"id_employee":     c.IdEmployee.Int64,
	})
}

type EmployeeShort struct {
	Firstname  sql.NullString `db:"employee_firstname" json:"employee_firstname"`
	Middlename sql.NullString `db:"employee_middlename" json:"employee_middlename"`
	Lastname   sql.NullString `db:"employee_lastname" json:"employee_lastname"`
	SkudCard   sql.NullString `db:"employee_skud_card" json:"employee_skud_card"`
}

type StudentShort struct {
	Firstname  sql.NullString `db:"student_firstname" json:"student_firstname"`
	Middlename sql.NullString `db:"student_middlename" json:"student_middlename"`
	Lastname   sql.NullString `db:"student_lastname" json:"student_lastname"`
	SkudCard   sql.NullString `db:"student_skud_card" json:"student_skud_card"`
}

type DatabaseMovement struct {
	Movement
	EmployeeShort
	StudentShort
}

func (c DatabaseMovement) MarshalJSON() ([]byte, error) {
	student := make(map[string]any)
	employee := make(map[string]any)
	movement := make(map[string]any)

	movement = map[string]any{
		"id":              c.Movement.Id,
		"id_building":     c.Movement.IdBuilding,
		"id_event":        c.Movement.IdEvent,
		"building_name":   c.Movement.BuildingName,
		"event_name":      c.Movement.EventName,
		"event_timestamp": c.Movement.EventTimestamp,
		"id_student":      c.Movement.IdStudent.Int64,
		"id_employee":     c.Movement.IdEmployee.Int64,
	}

	if c.Movement.IdStudent.Valid {
		student = map[string]any{
			"id":         c.Movement.IdStudent.Int64,
			"firstname":  c.StudentShort.Firstname.String,
			"middlename": c.StudentShort.Middlename.String,
			"lastname":   c.StudentShort.Lastname.String,
			"skud_card":  c.StudentShort.SkudCard.String,
		}
		return json.Marshal(map[string]map[string]any{
			"user":     student,
			"movement": movement,
		})
		// if c.Movement.IdEmployee.Valid
	} else {
		employee = map[string]any{
			"id":         c.Movement.IdEmployee.Int64,
			"firstname":  c.EmployeeShort.Firstname.String,
			"middlename": c.EmployeeShort.Middlename.String,
			"lastname":   c.EmployeeShort.Lastname.String,
			"skud_card":  c.EmployeeShort.SkudCard.String,
		}
		return json.Marshal(map[string]map[string]any{
			"user":     employee,
			"movement": movement,
		})
	}
}
