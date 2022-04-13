package main

import (
	"log"

	"github.com/AcuVuz/barriers-server/controllers"
	"github.com/AcuVuz/barriers-server/models"
	"github.com/AcuVuz/barriers-server/services"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sqlx.Connect("postgres", "host=db-jmu user=dbjmu password=Afgihn215zxdg dbname=jmu sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	usersController := &controllers.UsersController{
		UsersService: &services.UsersService{
			Students:  &models.StudentModel{DB: db},
			Employee:  &models.EmployeeModel{DB: db},
			Movements: &models.MovementModel{DB: db},
		},
	}

	app := gin.Default()

	usersRouter := app.Group("/users")
	{
		usersRouter.GET("/", usersController.Get)
		usersRouter.POST("/action", usersController.MovementAction)
	}

	app.Run(":8081")

}
