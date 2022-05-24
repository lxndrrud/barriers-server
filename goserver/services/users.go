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
		Get(id int64) (classes.DBStudentPersonalInfo, error)
		GetBySkudCard(SkudCard string) (classes.Student, error)
		GetGroupsInfo(IdStudent int64) ([]classes.DBStudentGroupInfo, error)
	}

	employee interface {
		Get(id int64) (classes.DBEmployeePersonalInfo, error)
		GetBySkudCard(SkudCard string) (classes.Employee, error)
		GetPositionsInfo(IdEmployee int64) ([]classes.DBEmployeePositionInfo, error)
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

func (s UsersService) GetStudentInfo(IdStudent int64) (classes.JSONStudentPersonalInfo,
	[]classes.JSONStudentGroupInfo, *classes.CustomError) {
	personal := classes.JSONStudentPersonalInfo{}
	groups := make([]classes.JSONStudentGroupInfo, 0)

	personalChan := make(chan classes.DBStudentPersonalInfo)
	groupsChan := make(chan []classes.DBStudentGroupInfo)
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

	personalInfo := <-personalChan
	groupsInfo := <-groupsChan
	err := <-errChan

	if err != nil {
		return personal, groups, &classes.CustomError{
			Code: 500,
			Text: err.Error(),
		}
	}

	personal = classes.CreateJSONStudentPersonalInfo(&personalInfo)
	groups = classes.CreateJSONStudentGroupInfo(groupsInfo)

	return personal, groups, nil
}

func (s UsersService) GetEmployeeInfo(IdEmployee int64) (classes.JSONEmployeePersonalInfo,
	[]classes.JSONEmployeePositionInfo, *classes.CustomError) {

	personal := classes.JSONEmployeePersonalInfo{}
	positions := make([]classes.JSONEmployeePositionInfo, 0)

	personalChan := make(chan classes.DBEmployeePersonalInfo)
	positionsChan := make(chan []classes.DBEmployeePositionInfo)
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

	EmployeePersonalInfo := <-personalChan
	EmployeePositionsInfo := <-positionsChan
	err := <-errChan

	if err != nil {
		return personal, positions, &classes.CustomError{
			Code: 500,
			Text: err.Error(),
		}
	}

	personal = classes.CreateJSONEmployeePersonalInfo(&EmployeePersonalInfo)
	positions = classes.CreateJSONEmployeePositionInfo(EmployeePositionsInfo)

	return personal, positions, nil
}
