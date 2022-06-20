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
		GetBySkudCard(SkudCard string) (classes.DBUser, *classes.CustomError)
	}
	movements interface {
		GetMovements(from time.Time, to time.Time) ([]classes.DatabaseMovement, error)
		GetMovementsForBuilding(idBuilding int64,
			from time.Time, to time.Time) ([]classes.DatabaseMovement, error)
		Insert(trx *sqlx.Tx, idBuilding, idEvent, idEmployee, idStudent int64) (int64, error)
		GetMovementsForUser(idStudent, idEmployee int64, from time.Time, to time.Time) ([]classes.Movement, error)
		GetMovementsForUserBuilding(idBuilding, idStudent, idEmployee int64,
			from time.Time, to time.Time) ([]classes.Movement, error)
	}
}

func CreateMovementsService(db *sqlx.DB) *MovementsService {
	return &MovementsService{
		db:           db,
		usersService: CreateUsersService(db),
		movements:    &models.MovementModel{DB: db},
	}
}

func (s MovementsService) MovementAction(idBuilding int64, event string,
	skudCard string) *classes.CustomError {
	user, err := s.usersService.GetBySkudCard(skudCard)
	if err != nil {
		return err
	}

	/*
		if user.Id == 0 {
			return &classes.CustomError{
				Text: "Пользователь не найден!",
				Code: http.StatusNotFound,
			}
		}
	*/

	var idEvent int64
	if event == "enter" {
		idEvent = 1
	} else if event == "exit" {
		idEvent = 2
	} else if event == "fail" {
		idEvent = 3
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

	// Обработчик Гостя - пользователя без карты
	if user.Id == -1 {
		_, err := s.movements.Insert(trx, idBuilding, idEvent, 0, 0)
		if err != nil {
			errTrx = trx.Rollback()
			if errTrx != nil {
				return &classes.CustomError{
					Text: errTrx.Error(),
					Code: http.StatusInternalServerError,
				}
			}
			return &classes.CustomError{
				Text: err.Error(),
				Code: http.StatusInternalServerError,
			}
		}
		// Обработчик студента
	} else if user.Type == "Студент" {
		_, err := s.movements.Insert(trx, idBuilding, idEvent, 0, user.Id)
		if err != nil {
			errTrx = trx.Rollback()
			if errTrx != nil {
				return &classes.CustomError{
					Text: errTrx.Error(),
					Code: http.StatusInternalServerError,
				}
			}
			return &classes.CustomError{
				Text: err.Error(),
				Code: http.StatusInternalServerError,
			}
		}
		// Обработчик сотрудника
	} else if user.Type == "Сотрудник" {
		_, err := s.movements.Insert(trx, idBuilding, idEvent, user.Id, 0)
		if err != nil {
			errTrx = trx.Rollback()
			if errTrx != nil {
				return &classes.CustomError{
					Text: errTrx.Error(),
					Code: http.StatusInternalServerError,
				}
			}
			return &classes.CustomError{
				Text: err.Error(),
				Code: http.StatusInternalServerError,
			}
		}
		// Обработчик неправильного типа человека
	} else {
		errTrx = trx.Rollback()
		if errTrx != nil {
			return &classes.CustomError{
				Text: errTrx.Error(),
				Code: http.StatusInternalServerError,
			}
		}
		return &classes.CustomError{
			Text: "Не определён тип человека!",
			Code: http.StatusInternalServerError,
		}

	}

	errTrx = trx.Commit()
	if errTrx != nil {
		return &classes.CustomError{
			Text: errTrx.Error(),
			Code: http.StatusInternalServerError,
		}
	}
	return nil
}

func (s MovementsService) GetMovements(idBuilding int64, from string, to string) ([]classes.JSONMovement, *classes.CustomError) {
	dateUtil := &utils.Dates{}
	var movementsJSON []classes.JSONMovement

	now := time.Now()
	parsedFrom := dateUtil.ParseWithDefault(from, time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()))
	parsedTo := dateUtil.ParseWithDefault(to, now.AddDate(0, 0, 1))

	dateUtil = nil

	var movements []classes.DatabaseMovement
	var err error

	if idBuilding == 0 {
		movements, err = s.movements.GetMovements(parsedFrom, parsedTo)
		if err != nil {
			return movementsJSON, &classes.CustomError{
				Text: "Внутренняя ошибка сервера при поиске перемещений!",
				Code: http.StatusInternalServerError,
			}
		}
	} else {
		movements, err = s.movements.GetMovementsForBuilding(idBuilding, parsedFrom, parsedTo)
		if err != nil {
			return movementsJSON, &classes.CustomError{
				Text: "Внутренняя ошибка сервера при поиске перемещений!",
				Code: http.StatusInternalServerError,
			}
		}
	}

	for _, movement := range movements {
		toAppend := classes.CreateJSONMovementFromDatabaseMovement(&movement)
		if toAppend.Id != 0 {
			movementsJSON = append(movementsJSON, classes.CreateJSONMovementFromDatabaseMovement(&movement))
		}
	}

	return movementsJSON, nil
}

func (s MovementsService) GetMovementsForUser(idBuilding, idEmployee, idStudent int64, from, to string) ([]classes.MovementJSON, *classes.CustomError) {
	movements := make([]classes.MovementJSON, 0)

	var parsedFrom time.Time
	var parsedTo time.Time

	dateUtil := &utils.Dates{}

	now := time.Now()
	parsedFrom = dateUtil.ParseWithDefault(from, time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()))
	parsedTo = dateUtil.ParseWithDefault(to, now.AddDate(0, 0, 1))

	dateUtil = nil

	var DBmovements []classes.Movement
	var err error

	if idBuilding == 0 {
		DBmovements, err = s.movements.GetMovementsForUser(idStudent, idEmployee, parsedFrom, parsedTo)
		if err != nil {
			return movements, &classes.CustomError{
				Code: 500,
				Text: err.Error(),
			}
		}
	} else {
		DBmovements, err = s.movements.GetMovementsForUserBuilding(idBuilding, idStudent, idEmployee,
			parsedFrom, parsedTo)
		if err != nil {
			return movements, &classes.CustomError{
				Code: 500,
				Text: err.Error(),
			}
		}
	}

	for _, movement := range DBmovements {
		movements = append(movements, classes.CreateJSONFromMovement(&movement))
	}

	return movements, nil
}
