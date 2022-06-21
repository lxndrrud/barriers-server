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

	buildingsRouter := app.Group("/buildings")
	SetupBuildingsRouter(buildingsRouter, db)

	return app
}
