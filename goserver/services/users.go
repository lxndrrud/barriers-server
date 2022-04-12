package services

import (
	"github.com/AcuVuz/barriers-server/models"
)

type UsersService struct {
	Students interface {
		GetBySkudCard(SkudCard string) (models.Student, error)
	}

	Persons interface {
		GetBySkudCard(SkudCard string) (models.Person, error)
	}
}

func (s UsersService) GetBySkudCard(SkudCard string) (models.Student, models.Person, error) {
	studentChan := make(chan models.Student)
	personChan := make(chan models.Person)
	errChan := make(chan error)

	go func() {
		value, err := s.Students.GetBySkudCard(SkudCard)
		studentChan <- value
		errChan <- err
	}()

	go func() {
		value, err := s.Persons.GetBySkudCard(SkudCard)
		personChan <- value
		errChan <- err
	}()

	student := <-studentChan
	person := <-personChan
	err := <-errChan

	if err != nil {
		return models.Student{}, models.Person{}, err
	}

	return student, person, nil
}
