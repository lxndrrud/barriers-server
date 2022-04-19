package controllers

import (
	"net/http"
	"strconv"

	"github.com/AcuVuz/barriers-server/interfaces"
	"github.com/AcuVuz/barriers-server/models"
	"github.com/AcuVuz/barriers-server/services"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type UsersController struct {
	UsersService interface {
		GetBySkudCard(SkudCard string) (models.Student, models.Employee, error)
		MovementAction(idBuilding int64, event string, skudCard string) *interfaces.CustomError
		GetMovements(from string, to string) ([]models.Movement, *interfaces.CustomError)
		GetMovementsForEmployee(idEmployee int64, from string, to string) ([]models.EmployeeMovement, *interfaces.CustomError)
		GetMovementsForStudent(idStudent int64, from string, to string) ([]models.StudentMovement, *interfaces.CustomError)
	}
}

func CreateUsersController(db *sqlx.DB) *UsersController {
	return &UsersController{
		UsersService: services.CreateUsersService(db),
	}
}

func (c UsersController) Get(ctx *gin.Context) {
	student, employee, err := c.UsersService.GetBySkudCard(ctx.Query("skud_card"))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	if student.Id != 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"student":  student,
			"employee": nil,
		})
	} else if employee.Id != 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"student":  nil,
			"employee": employee,
		})
	} else {
		ctx.JSON(http.StatusNotFound, gin.H{
			"student":  nil,
			"employee": nil,
		})
	}

}

func (c UsersController) MovementAction(ctx *gin.Context) {
	idBuilding, err := strconv.ParseInt(ctx.PostForm("id_building"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Возникла ошибка с определением идентификатора здания!",
		})
		return
	}

	event := ctx.PostForm("event")
	if len(event) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Не указан тип события!",
		})
		return
	}

	errService := c.UsersService.MovementAction(idBuilding, event, ctx.PostForm("skud_card"))
	if errService != nil {
		ctx.JSON(errService.Code, gin.H{
			"error": errService.Text,
		})
		errService = nil
		return
	}

	ctx.Status(http.StatusCreated)
}

func (c UsersController) GetMovements(ctx *gin.Context) {
	from := ctx.Query("from")
	to := ctx.Query("to")

	movementsQuery, errService := c.UsersService.GetMovements(from, to)
	if errService != nil {
		ctx.JSON(errService.Code, gin.H{
			"error": errService.Text,
		})
		errService = nil
		return
	}

	ctx.JSON(http.StatusOK, movementsQuery)

}

func (c UsersController) GetMovementsForEmployee(ctx *gin.Context) {
	from := ctx.Query("from")
	to := ctx.Query("to")
	idEmployee, err := strconv.ParseInt(ctx.Query("id_employee"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Ошибка при определении идентификатора работника!",
		})
		return
	}

	movementsQuery, errService := c.UsersService.GetMovementsForEmployee(idEmployee, from, to)
	if errService != nil {
		ctx.JSON(errService.Code, gin.H{
			"error": errService.Text,
		})
		errService = nil
		return
	}

	ctx.JSON(http.StatusOK, movementsQuery)
}

func (c UsersController) GetMovementsForStudent(ctx *gin.Context) {
	from := ctx.Query("from")
	to := ctx.Query("to")
	idStudent, err := strconv.ParseInt(ctx.Query("id_student"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Ошибка при определении идентификатора студента!",
		})
		return
	}

	movementsQuery, errService := c.UsersService.GetMovementsForStudent(idStudent, from, to)
	if errService != nil {
		ctx.JSON(errService.Code, gin.H{
			"error": errService.Text,
		})
		errService = nil
		return
	}

	ctx.JSON(http.StatusOK, movementsQuery)
}
