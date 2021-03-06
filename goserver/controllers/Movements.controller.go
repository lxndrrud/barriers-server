package controllers

import (
	"net/http"
	"strconv"

	"github.com/AcuVuz/barriers-server/classes"
	"github.com/AcuVuz/barriers-server/services"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type MovementsController struct {
	MovementsService interface {
		MovementAction(idBuilding int64, event string, skudCard string) *classes.CustomError
		GetMovements(idBuilding int64,
			from string, to string) ([]classes.DatabaseMovement, *classes.CustomError)
		//GetMovementsForEmployee(idEmployee int64, from string, to string) ([]classes.JSONEmployeeMovement, *classes.CustomError)
		//GetMovementsForStudent(idStudent int64, from string, to string) ([]classes.JSONStudentMovement, *classes.CustomError)
		GetMovementsForUser(idBuilding, idEmployee, idStudent int64,
			from, to string) ([]classes.Movement, *classes.CustomError)
	}
}

func CreateMovementsController(db *sqlx.DB) *MovementsController {
	return &MovementsController{
		MovementsService: services.CreateMovementsService(db),
	}
}

func (c MovementsController) MovementAction(ctx *gin.Context) {
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

	errService := c.MovementsService.MovementAction(idBuilding, event, ctx.PostForm("skud_card"))
	if errService != nil {
		ctx.JSON(errService.Code, gin.H{
			"error": errService.Text,
		})
		errService = nil
		return
	}

	ctx.Status(http.StatusCreated)
}

func (c MovementsController) GetMovements(ctx *gin.Context) {
	from := ctx.Query("from")
	to := ctx.Query("to")
	idBuilding, err := strconv.ParseInt(ctx.Query("id_building"), 10, 64)
	if err != nil {
		idBuilding = 0
	}

	movementsQuery, errService := c.MovementsService.GetMovements(idBuilding, from, to)
	if errService != nil {
		ctx.JSON(errService.Code, gin.H{
			"error": errService.Text,
		})
		errService = nil
		return
	}

	if len(movementsQuery) == 0 {
		ctx.JSON(http.StatusOK, []classes.Movement{})
		return
	}

	ctx.JSON(http.StatusOK, movementsQuery)

}

/*
func (c MovementsController) GetMovementsForEmployee(ctx *gin.Context) {
	from := ctx.Query("from")
	to := ctx.Query("to")
	idEmployee, err := strconv.ParseInt(ctx.Query("id_employee"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Ошибка при определении идентификатора работника!",
		})
		return
	}

	movementsQuery, errService := c.MovementsService.GetMovementsForEmployee(idEmployee, from, to)
	if errService != nil {
		ctx.JSON(errService.Code, gin.H{
			"error": errService.Text,
		})
		errService = nil
		return
	}

	ctx.JSON(http.StatusOK, movementsQuery)
}

func (c MovementsController) GetMovementsForStudent(ctx *gin.Context) {
	from := ctx.Query("from")
	to := ctx.Query("to")
	idStudent, err := strconv.ParseInt(ctx.Query("id_student"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Ошибка при определении идентификатора студента!",
		})
		return
	}

	movementsQuery, errService := c.MovementsService.GetMovementsForStudent(idStudent, from, to)
	if errService != nil {
		ctx.JSON(errService.Code, gin.H{
			"error": errService.Text,
		})
		errService = nil
		return
	}

	ctx.JSON(http.StatusOK, movementsQuery)
}
*/

func (c MovementsController) GetMovementsForUser(ctx *gin.Context) {
	from := ctx.Query("from")
	to := ctx.Query("to")
	idBuilding, err := strconv.ParseInt(ctx.Query("id_building"), 10, 64)
	if err != nil {
		idBuilding = 0
	}

	idStudent, err := strconv.ParseInt(ctx.Query("id_student"), 10, 64)
	if err != nil {
		idStudent = 0
	}

	idEmployee, err := strconv.ParseInt(ctx.Query("id_employee"), 10, 64)
	if err != nil {
		idEmployee = 0
	}

	movements, errService := c.MovementsService.
		GetMovementsForUser(idBuilding, idEmployee, idStudent, from, to)
	if errService != nil {
		ctx.JSON(errService.Code, errService.Text)
		errService = nil
		return
	}

	ctx.JSON(http.StatusOK, movements)
}
