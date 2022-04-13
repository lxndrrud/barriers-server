package services

import (
	"net/http"

	"github.com/AcuVuz/barriers-server/interfaces"
	"github.com/AcuVuz/barriers-server/models"
)

type UsersService struct {
	Students *models.StudentModel

	Employee *models.EmployeeModel

	Movements *models.MovementModel
}

func (s UsersService) GetBySkudCard(SkudCard string) (models.Student, models.Employee, error) {
	studentChan := make(chan models.Student)
	personChan := make(chan models.Employee)
	errChan := make(chan error)

	go func() {
		value, err := s.Students.GetBySkudCard(SkudCard)
		studentChan <- value
		errChan <- err
	}()

	go func() {
		value, err := s.Employee.GetBySkudCard(SkudCard)
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

func (s UsersService) MovementAction(idBuilding int64, idEvent int64, skudCard string) *interfaces.CustomError {
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

	if student.Id != 0 {
		_, err := s.Movements.InsertForStudent(idBuilding, idEvent, student.Id)
		if err != nil {
			return &interfaces.CustomError{
				Text: err.Error(),
				Code: http.StatusInternalServerError,
			}
		}
	} else if employee.Id != 0 {
		_, err := s.Movements.InsertForEmployee(idBuilding, idEvent, employee.Id)
		if err != nil {
			return &interfaces.CustomError{
				Text: err.Error(),
				Code: http.StatusInternalServerError,
			}
		}
	}
	return nil
}

func (s UsersService) GetMovements(from string, to string) ([]models.Movement, error) {
	return []models.Movement{}, nil
}
