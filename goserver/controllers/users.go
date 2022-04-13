package controllers

import (
	"net/http"
	"strconv"

	"github.com/AcuVuz/barriers-server/services"
	"github.com/gin-gonic/gin"
)

type UsersController struct {
	UsersService *services.UsersService
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

	idEvent, err := strconv.ParseInt(ctx.PostForm("id_event"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Возникла ошибка с определением идентификатора события перемещения!",
		})
		return
	}

	errService := c.UsersService.MovementAction(idBuilding, idEvent, ctx.PostForm("skud_card"))
	if errService != nil {
		ctx.JSON(errService.Code, gin.H{
			"error": errService.Error(),
		})
		return
	}

	ctx.Status(http.StatusCreated)
}

func (c UsersController) GetMovements(ctx *gin.Context) {

}
