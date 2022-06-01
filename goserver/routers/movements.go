package routers

import (
	"github.com/AcuVuz/barriers-server/controllers"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func SetupMovementsRouter(usersRouter *gin.RouterGroup, db *sqlx.DB) {
	movementsController := controllers.CreateMovementsController(db)

	usersRouter.POST("/action", movementsController.MovementAction)
	usersRouter.GET("/", movementsController.GetMovements)
	/*
		usersRouter.GET("/employee", movementsController.GetMovementsForEmployee)
		usersRouter.GET("/student", movementsController.GetMovementsForStudent)
	*/
	usersRouter.GET("/user", movementsController.GetMovementsForUser)

}
