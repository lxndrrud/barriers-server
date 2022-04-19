package routers

import (
	"github.com/AcuVuz/barriers-server/controllers"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func SetupUsersRouter(usersRouter *gin.RouterGroup, db *sqlx.DB) {
	usersController := controllers.CreateUsersController(db)

	usersRouter.GET("/", usersController.Get)
	usersRouter.POST("/action", usersController.MovementAction)
	usersRouter.GET("/movements", usersController.GetMovements)
	usersRouter.GET("/movements/employee", usersController.GetMovementsForEmployee)
	usersRouter.GET("/movements/student", usersController.GetMovementsForStudent)

}
