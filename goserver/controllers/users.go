package controllers

import (
	"net/http"

	"github.com/AcuVuz/barriers-server/classes"
	"github.com/AcuVuz/barriers-server/services"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type UsersController struct {
	UsersService interface {
		GetBySkudCard(SkudCard string) (classes.Student, classes.Employee, *classes.CustomError)
	}
}

func CreateUsersController(db *sqlx.DB) *UsersController {
	return &UsersController{
		UsersService: services.CreateUsersService(db),
	}
}

func (c UsersController) GetBySkudCard(ctx *gin.Context) {
	student, employee, err := c.UsersService.GetBySkudCard(ctx.Query("skud_card"))
	if err != nil {
		ctx.JSON(err.Code, gin.H{
			"error": err.Error(),
		})
		return
	}
	if student.Id != 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"student":  student.User,
			"employee": nil,
		})
	} else if employee.Id != 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"student":  nil,
			"employee": employee.User,
		})
	} else {
		ctx.JSON(http.StatusNotFound, gin.H{
			"student":  nil,
			"employee": nil,
		})
	}

}
