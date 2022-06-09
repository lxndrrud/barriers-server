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
		Get(id int64) (classes.DBUser, error)
		GetBySkudCard(SkudCard string) (classes.DBUser, error)
		GetGroupsInfo(IdStudent int64) ([]classes.DBStudentGroupInfo1, error)
	}

	employee interface {
		Get(id int64) (classes.DBUser, error)
		GetBySkudCard(SkudCard string) (classes.DBUser, error)
		GetPositionsInfo(IdEmployee int64) ([]classes.DBEmployeePositionInfo1, error)
	}
}

func CreateUsersService(db *sqlx.DB) *UsersService {
	return &UsersService{
		db:       db,
		students: &models.StudentModel{DB: db},
		employee: &models.EmployeeModel{DB: db},
	}
}

func (s UsersService) GetBySkudCard(SkudCard string) (classes.DBUser, *classes.CustomError) {
	studentChan := make(chan classes.DBUser)
	employeeChan := make(chan classes.DBUser)
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
		return classes.DBUser{}, &classes.CustomError{
			Text: err.Error(),
			Code: http.StatusInternalServerError,
		}
	}

	if student.Id != 0 {
		return student, nil
	} else if employee.Id != 0 {
		return student, nil
	} else {
		return classes.DBUser{}, nil
	}

}

func (s UsersService) GetStudentInfo(IdStudent int64) (classes.JSONStudent, *classes.CustomError) {
	personal := classes.DBUser{}
	groups := make([]classes.DBStudentGroupInfo1, 0)

	personalChan := make(chan classes.DBUser)
	groupsChan := make(chan []classes.DBStudentGroupInfo1)
	errChan := make(chan error)

	go func() {
		value, err := s.students.Get(IdStudent)
		personalChan <- value
		errChan <- err
	}()

	go func() {
		value, err := s.students.GetGroupsInfo(IdStudent)
		groupsChan <- value
		errChan <- err
	}()

	personal = <-personalChan
	groups = <-groupsChan
	err := <-errChan

	if err != nil {
		return classes.JSONStudent{}, &classes.CustomError{
			Code: 500,
			Text: err.Error(),
		}
	}

	return classes.JSONStudent{
		Student: personal,
		Groups:  groups,
	}, nil
}

func (s UsersService) GetEmployeeInfo(IdEmployee int64) (classes.JSONEmployee, *classes.CustomError) {

	personal := classes.DBUser{}
	positions := make([]classes.DBEmployeePositionInfo1, 0)

	personalChan := make(chan classes.DBUser)
	positionsChan := make(chan []classes.DBEmployeePositionInfo1)
	errChan := make(chan error)

	go func() {
		value, err := s.employee.Get(IdEmployee)
		personalChan <- value
		errChan <- err
	}()

	go func() {
		value, err := s.employee.GetPositionsInfo(IdEmployee)
		positionsChan <- value
		errChan <- err
	}()

	personal = <-personalChan
	positions = <-positionsChan
	err := <-errChan

	if err != nil {
		return classes.JSONEmployee{}, &classes.CustomError{
			Code: 500,
			Text: err.Error(),
		}
	}

	return classes.JSONEmployee{
		Employee:  personal,
		Positions: positions,
	}, nil
}
