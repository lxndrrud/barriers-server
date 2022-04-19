package services

import (
	"net/http"
	"time"

	"github.com/AcuVuz/barriers-server/interfaces"
	"github.com/AcuVuz/barriers-server/models"
	"github.com/AcuVuz/barriers-server/utils"
	"github.com/jmoiron/sqlx"
)

type UsersService struct {
	db *sqlx.DB

	students interface {
		Get(id int64) (models.Student, error)
		GetBySkudCard(SkudCard string) (models.Student, error)
	}

	employee interface {
		Get(id int64) (models.Employee, error)
		GetBySkudCard(SkudCard string) (models.Employee, error)
	}

	movements interface {
		GetMovements(from time.Time, to time.Time) ([]models.Movement, error)
		InsertForStudent(trx *sqlx.Tx, idBuilding int64, idEvent int64, idStudent int64) (int64, error)
		InsertForEmployee(trx *sqlx.Tx, idBuilding int64, idEvent int64, idEmployee int64) (int64, error)
		GetMovementsForEmployee(idEmployee int64, from time.Time, to time.Time) ([]models.EmployeeMovement, error)
		GetMovementsForStudent(idStudent int64, from time.Time, to time.Time) ([]models.StudentMovement, error)
	}
}

func CreateUsersService(db *sqlx.DB) *UsersService {
	return &UsersService{
		db:        db,
		students:  &models.StudentModel{DB: db},
		employee:  &models.EmployeeModel{DB: db},
		movements: &models.MovementModel{DB: db},
	}
}

func (s UsersService) GetBySkudCard(SkudCard string) (models.Student, models.Employee, error) {
	studentChan := make(chan models.Student)
	personChan := make(chan models.Employee)
	errChan := make(chan error)

	go func() {
		value, err := s.students.GetBySkudCard(SkudCard)
		studentChan <- value
		errChan <- err
	}()

	go func() {
		value, err := s.employee.GetBySkudCard(SkudCard)
		personChan <- value
		errChan <- err
	}()

	student := <-studentChan
	person := <-personChan
	err := <-errChan

	if err != nil {
		return models.Student{}, models.Employee{}, err
	}

	return student, person, nil
}

func (s UsersService) MovementAction(idBuilding int64, event string, skudCard string) *interfaces.CustomError {
	student, employee, err := s.GetBySkudCard(skudCard)
	if err != nil {
		return &interfaces.CustomError{
			Text: err.Error(),
			Code: http.StatusInternalServerError,
		}
	}

	if student.Id == 0 && employee.Id == 0 {
		return &interfaces.CustomError{
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
		return &interfaces.CustomError{
			Text: "Неверно указан тип события!",
			Code: http.StatusBadRequest,
		}
	}

	trx, err := s.db.Beginx()
	if err != nil {
		return &interfaces.CustomError{
			Text: err.Error(),
			Code: http.StatusInternalServerError,
		}
	}

	if student.Id != 0 {
		_, err := s.movements.InsertForStudent(trx, idBuilding, idEvent, student.Id)
		if err != nil {
			return &interfaces.CustomError{
				Text: err.Error(),
				Code: http.StatusInternalServerError,
			}
		}
	} else if employee.Id != 0 {
		_, err := s.movements.InsertForEmployee(trx, idBuilding, idEvent, employee.Id)
		if err != nil {
			return &interfaces.CustomError{
				Text: err.Error(),
				Code: http.StatusInternalServerError,
			}
		}
	}

	err = trx.Commit()
	if err != nil {
		return &interfaces.CustomError{
			Text: err.Error(),
			Code: http.StatusInternalServerError,
		}
	}
	return nil
}

func (s UsersService) GetMovements(from string, to string) ([]models.Movement, *interfaces.CustomError) {
	dateUtil := &utils.Dates{}

	now := time.Now()
	parsedFrom := dateUtil.ParseWithDefault(from, time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()))
	parsedTo := dateUtil.ParseWithDefault(to, now.AddDate(0, 0, 1))

	dateUtil = nil

	movements, err := s.movements.GetMovements(parsedFrom, parsedTo)
	if err != nil {
		return movements, &interfaces.CustomError{
			Text: "Внутренняя ошибка сервера при поиске перемещений!",
			Code: http.StatusInternalServerError,
		}
	}

	return movements, nil
}

func (s UsersService) GetMovementsForEmployee(idEmployee int64, from string, to string) ([]models.EmployeeMovement, *interfaces.CustomError) {
	var movements []models.EmployeeMovement

	dateUtil := &utils.Dates{}

	now := time.Now()
	parsedFrom := dateUtil.ParseWithDefault(from, time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()))
	parsedTo := dateUtil.ParseWithDefault(to, now.AddDate(0, 0, 1))

	dateUtil = nil

	movements, err := s.movements.GetMovementsForEmployee(idEmployee, parsedFrom, parsedTo)
	if err != nil {
		return movements, &interfaces.CustomError{
			Text: "Внутренняя ошибка сервера при поиске перемещений для работника!",
			Code: http.StatusInternalServerError,
		}
	}
	return movements, nil
}

func (s UsersService) GetMovementsForStudent(idStudent int64, from string, to string) ([]models.StudentMovement, *interfaces.CustomError) {
	var movements []models.StudentMovement

	var parsedFrom time.Time
	var parsedTo time.Time

	dateUtil := &utils.Dates{}

	now := time.Now()
	parsedFrom = dateUtil.ParseWithDefault(from, time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()))
	parsedTo = dateUtil.ParseWithDefault(to, now.AddDate(0, 0, 1))

	dateUtil = nil

	movements, err := s.movements.GetMovementsForStudent(idStudent, parsedFrom, parsedTo)
	if err != nil {
		return movements, &interfaces.CustomError{
			Text: "Внутренняя ошибка сервера при поиске перемещений для студента!",
			Code: http.StatusInternalServerError,
		}
	}
	return movements, nil
}
