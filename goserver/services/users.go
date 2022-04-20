package services

import (
	"net/http"

	"github.com/AcuVuz/barriers-server/classes"
	"github.com/AcuVuz/barriers-server/models"
	"github.com/jmoiron/sqlx"
)

type UsersService struct {
	db *sqlx.DB

	students interface {
		Get(id int64) (classes.Student, error)
		GetBySkudCard(SkudCard string) (classes.Student, error)
	}

	employee interface {
		Get(id int64) (classes.Employee, error)
		GetBySkudCard(SkudCard string) (classes.Employee, error)
	}
}

func CreateUsersService(db *sqlx.DB) *UsersService {
	return &UsersService{
		db:       db,
		students: &models.StudentModel{DB: db},
		employee: &models.EmployeeModel{DB: db},
	}
}

func (s UsersService) GetBySkudCard(SkudCard string) (classes.Student, classes.Employee, *classes.CustomError) {
	studentChan := make(chan classes.Student)
	personChan := make(chan classes.Employee)
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
		return classes.Student{}, classes.Employee{}, &classes.CustomError{
			Text: err.Error(),
			Code: http.StatusInternalServerError,
		}
	}

	return student, person, nil
}
