package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func GetApp(db *sqlx.DB) *gin.Engine {

	app := gin.Default()

	usersRouter := app.Group("/users")
	SetupUsersRouter(usersRouter, db)

	movementsRouter := app.Group("/movements")
	SetupMovementsRouter(movementsRouter, db)

	return app
}
