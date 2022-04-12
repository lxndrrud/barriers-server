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
			Students: &models.StudentModel{DB: db},
			Persons:  &models.PersonModel{DB: db},
		},
	}

	app := gin.Default()

	app.GET("/", usersController.Get)

	app.Run(":8081")

}
