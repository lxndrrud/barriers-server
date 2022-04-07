package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/AcuVuz/barriers-server/models"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Env struct {
	students interface {
		GetBySkudCard(SkudCard string) (*models.Student, error)
	}

	persons interface {
		GetBySkudCard(SkudCard string) (*models.Person, error)
	}
}

func (e Env) Get(ctx *gin.Context) {
	student, err := e.students.GetBySkudCard(ctx.Query("skud_card"))
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	if student != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"student": student,
			"person":  nil,
		})
		return
	}
	person, err := e.persons.GetBySkudCard(ctx.Query("skud_card"))
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	if person != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"student": nil,
			"person":  person,
		})
		return
	}
	ctx.JSON(http.StatusNotFound, gin.H{
		"error": err.Error(),
	})
}

func main() {
	db, err := sqlx.Connect("postgres", "host=db-jmu user=dbjmu password=Afgihn215zxdg dbname=jmu sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	env := &Env{
		students: models.StudentModel{DB: db},
		persons:  models.PersonModel{DB: db},
	}

	app := gin.Default()

	app.GET("/", env.Get)

	app.Run(":8081")

}
