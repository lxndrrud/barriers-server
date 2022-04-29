package controllers

import (
	"net/http"

	"github.com/AcuVuz/barriers-server/classes"
	"github.com/AcuVuz/barriers-server/services"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type UsersController struct {
	UsersService interface {
		GetBySkudCard(SkudCard string) (classes.UserJSON, *classes.CustomError)
	}
}

func CreateUsersController(db *sqlx.DB) *UsersController {
	return &UsersController{
		UsersService: services.CreateUsersService(db),
	}
}

func (c UsersController) GetBySkudCard(ctx *gin.Context) {
	user, err := c.UsersService.GetBySkudCard(ctx.Query("skud_card"))
	if err != nil {
		ctx.JSON(err.Code, gin.H{
			"error": err.Error(),
		})
		return
	}
	if user.Id != 0 {
		ctx.JSON(http.StatusOK, user)
	} else {
		ctx.JSON(http.StatusNotFound, user)
	}

}
