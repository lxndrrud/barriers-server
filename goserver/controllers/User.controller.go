package controllers

import (
	"net/http"
	"strconv"

	"github.com/AcuVuz/barriers-server/classes"
	"github.com/AcuVuz/barriers-server/services"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type UsersController struct {
	UsersService interface {
		GetBySkudCard(SkudCard string) (classes.DBUser, *classes.CustomError)
		GetEmployeeInfo(IdEmployee int64) (classes.JSONEmployee,
			*classes.CustomError)
		GetStudentInfo(IdStudent int64) (classes.JSONStudent, *classes.CustomError)
	}
}

func CreateUsersController(db *sqlx.DB) *UsersController {
	return &UsersController{
		UsersService: services.CreateUsersService(db),
	}
}

func (c UsersController) GetBySkudCard(ctx *gin.Context) {
	user, err := c.UsersService.GetBySkudCard(ctx.Query("skud_card"))
	if err != nil {
		ctx.JSON(err.Code, gin.H{
			"error": err.Error(),
		})
		return
	}
	if user.Id != 0 {
		ctx.JSON(http.StatusOK, user)
	} else {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Пользователь не найден!",
		})
	}

}

func (c UsersController) GetEmployeeInfo(ctx *gin.Context) {
	idEmployee, err := strconv.ParseInt(ctx.Query("id_employee"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Ошибка при определении идентификатора работника!",
		})
		return
	}
	employee, errService := c.UsersService.GetEmployeeInfo(idEmployee)
	if errService != nil {
		ctx.JSON(errService.Code, errService.Text)
		return
	}

	ctx.JSON(http.StatusOK, employee)
}

func (c UsersController) GetStudentInfo(ctx *gin.Context) {
	idStudent, err := strconv.ParseInt(ctx.Query("id_student"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Ошибка при определении идентификатора работника!",
		})
		return
	}
	student, errService := c.UsersService.GetStudentInfo(idStudent)
	if errService != nil {
		ctx.JSON(errService.Code, errService.Text)
		return
	}

	ctx.JSON(http.StatusOK, student)
}
