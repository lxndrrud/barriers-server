package controllers

import (
	"net/http"

	"github.com/AcuVuz/barriers-server/models"
	"github.com/gin-gonic/gin"
)

type UsersController struct {
	UsersService interface {
		GetBySkudCard(SkudCard string) (models.Student, models.Person, error)
	}
}

func (c UsersController) Get(ctx *gin.Context) {
	student, person, err := c.UsersService.GetBySkudCard(ctx.Query("skud_card"))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	if student.Id != 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"student": student,
			"person":  nil,
		})
	} else if person.Id != 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"student": nil,
			"person":  person,
		})
	} else {
		ctx.JSON(http.StatusNotFound, gin.H{
			"student": nil,
			"person":  nil,
		})
	}

}
