package controllers

import (
	"net/http"

	"github.com/AcuVuz/barriers-server/classes"
	"github.com/AcuVuz/barriers-server/services"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type BuildingsController struct {
	buildingsService interface {
		GetAll() ([]classes.Building, *classes.CustomError)
	}
}

func CreateBuildingsController(db *sqlx.DB) *BuildingsController {
	return &BuildingsController{
		buildingsService: services.CreateBuildingsService(db),
	}
}

func (c BuildingsController) GetAll(ctx *gin.Context) {
	buildings, errService := c.buildingsService.GetAll()
	if errService != nil {
		ctx.JSON(errService.Code, gin.H{
			"error": errService.Text,
		})
		return
	}
	ctx.JSON(http.StatusOK, buildings)
}
