package main

import (
	"log"

	"github.com/AcuVuz/barriers-server/routers"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sqlx.Connect("postgres",
		"host=test user=test password=test dbname=test sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	app := routers.GetApp(db)

	app.Run(":8081")

}
