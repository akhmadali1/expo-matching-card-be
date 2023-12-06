package main

import (
	"log"
	_ "match_card/docs"
	"match_card/routes"

	"github.com/joho/godotenv"
)

func main() {

	eri := godotenv.Load() // 👈 load .env file
	if eri != nil {
		log.Fatal(eri)
		return
	}

	router := routes.SetupRoutes()

	router.Run(":8082")
}
