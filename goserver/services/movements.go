package services

import (
	"net/http"
	"time"

	"github.com/AcuVuz/barriers-server/classes"
	"github.com/AcuVuz/barriers-server/models"
	"github.com/AcuVuz/barriers-server/utils"
	"github.com/jmoiron/sqlx"
)

type MovementsService struct {
	db           *sqlx.DB
	usersService interface {
		GetBySkudCard(SkudCard string) (classes.Student, classes.Employee, *classes.CustomError)
	}
	movements interface {
		GetMovements(from time.Time, to time.Time) ([]classes.Movement, error)
		InsertForStudent(trx *sqlx.Tx, idBuilding int64, idEvent int64, idStudent int64) (int64, error)
		InsertForEmployee(trx *sqlx.Tx, idBuilding int64, idEvent int64, idEmployee int64) (int64, error)
		GetMovementsForEmployee(idEmployee int64, from time.Time, to time.Time) ([]classes.EmployeeMovement, error)
		GetMovementsForStudent(idStudent int64, from time.Time, to time.Time) ([]classes.StudentMovement, error)
	}
}

func CreateMovementsService(db *sqlx.DB) *MovementsService {
	return &MovementsService{
		db:           db,
		usersService: CreateUsersService(db),
		movements:    &models.MovementModel{DB: db},
	}
}

func (s MovementsService) MovementAction(idBuilding int64, event string, skudCard string) *classes.CustomError {
	student, employee, err := s.usersService.GetBySkudCard(skudCard)
	if err != nil {
		return err
	}

	if student.User.Id == 0 && employee.User.Id == 0 {
		return &classes.CustomError{
			Text: "Пользователь не найден!",
			Code: http.StatusNotFound,
		}
	}

	var idEvent int64
	if event == "enter" {
		idEvent = 1
	} else if event == "exit" {
		idEvent = 2
	} else {
		return &classes.CustomError{
			Text: "Неверно указан тип события!",
			Code: http.StatusBadRequest,
		}
	}

	trx, errTrx := s.db.Beginx()
	if errTrx != nil {
		return &classes.CustomError{
			Text: errTrx.Error(),
			Code: http.StatusInternalServerError,
		}
	}

	if student.User.Id != 0 {
		_, err := s.movements.InsertForStudent(trx, idBuilding, idEvent, student.User.Id)
		if err != nil {
			return &classes.CustomError{
				Text: err.Error(),
				Code: http.StatusInternalServerError,
			}
		}
	} else if employee.User.Id != 0 {
		_, err := s.movements.InsertForEmployee(trx, idBuilding, idEvent, employee.User.Id)
		if err != nil {
			return &classes.CustomError{
				Text: err.Error(),
				Code: http.StatusInternalServerError,
			}
		}
	}

	errTrx = trx.Commit()
	if err != nil {
		return &classes.CustomError{
			Text: errTrx.Error(),
			Code: http.StatusInternalServerError,
		}
	}
	return nil
}

func (s MovementsService) GetMovements(from string, to string) ([]classes.MovementJSON, *classes.CustomError) {
	dateUtil := &utils.Dates{}
	var movementsJSON []classes.MovementJSON

	now := time.Now()
	parsedFrom := dateUtil.ParseWithDefault(from, time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()))
	parsedTo := dateUtil.ParseWithDefault(to, now.AddDate(0, 0, 1))

	dateUtil = nil

	movements, err := s.movements.GetMovements(parsedFrom, parsedTo)
	if err != nil {
		return movementsJSON, &classes.CustomError{
			Text: "Внутренняя ошибка сервера при поиске перемещений!",
			Code: http.StatusInternalServerError,
		}
	}

	for _, movement := range movements {
		movementsJSON = append(movementsJSON, classes.MovementJSON{
			Id:             movement.Id,
			IdBuilding:     movement.IdBuilding,
			IdEvent:        movement.IdEvent,
			EventName:      movement.EventName,
			EventTimestamp: movement.EventTimestamp,
			IdStudent:      movement.IdStudent.Int64,
			IdEmployee:     movement.IdEmployee.Int64,
		})
	}

	return movementsJSON, nil
}

func (s MovementsService) GetMovementsForEmployee(idEmployee int64, from string, to string) ([]classes.EmployeeMovement, *classes.CustomError) {
	var movements []classes.EmployeeMovement

	dateUtil := &utils.Dates{}

	now := time.Now()
	parsedFrom := dateUtil.ParseWithDefault(from, time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()))
	parsedTo := dateUtil.ParseWithDefault(to, now.AddDate(0, 0, 1))

	dateUtil = nil

	movements, err := s.movements.GetMovementsForEmployee(idEmployee, parsedFrom, parsedTo)
	if err != nil {
		return movements, &classes.CustomError{
			Text: "Внутренняя ошибка сервера при поиске перемещений для работника!",
			Code: http.StatusInternalServerError,
		}
	}
	return movements, nil
}

func (s MovementsService) GetMovementsForStudent(idStudent int64, from string, to string) ([]classes.StudentMovement, *classes.CustomError) {
	var movements []classes.StudentMovement

	var parsedFrom time.Time
	var parsedTo time.Time

	dateUtil := &utils.Dates{}

	now := time.Now()
	parsedFrom = dateUtil.ParseWithDefault(from, time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()))
	parsedTo = dateUtil.ParseWithDefault(to, now.AddDate(0, 0, 1))

	dateUtil = nil

	movements, err := s.movements.GetMovementsForStudent(idStudent, parsedFrom, parsedTo)
	if err != nil {
		return movements, &classes.CustomError{
			Text: "Внутренняя ошибка сервера при поиске перемещений для студента!",
			Code: http.StatusInternalServerError,
		}
	}
	return movements, nil
}
