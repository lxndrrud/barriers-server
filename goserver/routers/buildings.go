package routers

import (
	"github.com/AcuVuz/barriers-server/controllers"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func SetupBuildingsRouter(buildingsRouter *gin.RouterGroup, db *sqlx.DB) {
	buildingsController := controllers.CreateBuildingsController(db)

	buildingsRouter.GET("/", buildingsController.GetAll)
}
