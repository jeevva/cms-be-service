package main

import (
	"jeevva/cms-be-service/internal/routers"
	"jeevva/cms-be-service/pkg"
	"log"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	database, err := pkg.Pgdb()
	if err != nil {
		log.Fatal(err)
	}
	router := routers.New(database)
	server := pkg.Server(router)

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
