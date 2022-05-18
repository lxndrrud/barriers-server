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

func (s UsersService) GetBySkudCard(SkudCard string) (classes.UserJSON, *classes.CustomError) {
	studentChan := make(chan classes.Student)
	employeeChan := make(chan classes.Employee)
	errChan := make(chan error)

	go func() {
		value, err := s.students.GetBySkudCard(SkudCard)
		studentChan <- value
		errChan <- err
	}()

	go func() {
		value, err := s.employee.GetBySkudCard(SkudCard)
		employeeChan <- value
		errChan <- err
	}()

	student := <-studentChan
	employee := <-employeeChan
	err := <-errChan

	if err != nil {
		return classes.UserJSON{}, &classes.CustomError{
			Text: err.Error(),
			Code: http.StatusInternalServerError,
		}
	}

	if student.Id != 0 {
		return classes.CreateUserJSONFromStudent(&student), nil
	} else if employee.Id != 0 {
		return classes.CreateUserJSONFromEmployee(&employee), nil
	} else {
		return classes.UserJSON{}, nil
	}

}

func (s UsersService) GetStudentInfo(IdStudent int64) {

}

func (s UsersService) GetEmployeeInfo(IdEmployee int64) {

}
