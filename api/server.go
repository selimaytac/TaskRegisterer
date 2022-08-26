package api

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/selimaytac/TaskRegisterer/api/controllers"
	"github.com/selimaytac/TaskRegisterer/api/models"
	"github.com/selimaytac/TaskRegisterer/api/seed"
	"log"
	"os"
)

var server = controllers.Server{}

func Run() {
	var err error

	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}

	server.Initialize(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))

	if server.DB.HasTable(&models.User{}) == false {
		seed.Load(server.DB)
	}

	server.Run(":8080")
}
