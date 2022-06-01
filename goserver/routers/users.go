package routers

import (
	"github.com/AcuVuz/barriers-server/controllers"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func SetupUsersRouter(usersRouter *gin.RouterGroup, db *sqlx.DB) {
	usersController := controllers.CreateUsersController(db)

	usersRouter.GET("/skudCard", usersController.GetBySkudCard)
	usersRouter.GET("/employee", usersController.GetEmployeeInfo)

	// TODO Доделать эту штуку для студента
	usersRouter.GET("/student", usersController.GetStudentInfo)
}
