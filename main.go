package main

import (
	"os"
	// "fmt"
	"log"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	a := App{}
	app_db_username := os.Getenv("APP_DB_USERNAME")
	if app_db_username == "" {
		log.Fatal("APP_DB_USERNAME not found in the environment")
	}
	app_db_password := os.Getenv("APP_DB_PASSWORD")
	if app_db_password == "" {
		log.Fatal("APP_DB_PASSWORD not found in the environment")
	}
	app_db_name := os.Getenv("APP_DB_NAME")
	if app_db_name == "" {
		log.Fatal("APP_DB_NAME not found in the environment")
	}
	a.Initialize(
		app_db_username,
		app_db_password,
		app_db_name,
	)

	a.Run(":8010")
}