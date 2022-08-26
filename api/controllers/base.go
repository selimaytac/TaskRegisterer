package controllers

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/selimaytac/TaskRegisterer/api/models"
	"log"
	"net/http"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (server *Server) Initialize(DbDriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {
	var err error

	if DbDriver == "mysql" {
		DbURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)
		server.DB, err = gorm.Open(DbDriver, DbURL)

		if err != nil {
			fmt.Printf("Cannot connect to %s database", DbDriver)
			log.Fatal("This is the error:", err)
		} else {
			fmt.Printf("Connected to %s database", DbDriver)
		}
	}

	if DbDriver == "postgres" {
		DbURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
		server.DB, err = gorm.Open(DbDriver, DbURL)

		if err != nil {
			fmt.Printf("Cannot connect to %s database", DbDriver)
			log.Fatal("This is the error:", err)
		} else {
			fmt.Printf("Connected to %s database", DbDriver)
		}
	}

	server.DB.Debug().AutoMigrate(&models.User{})

	server.Router = mux.NewRouter()

	server.InitializeRoutes()
}

func (server *Server) Run(add string) {
	fmt.Println("Listening to port 8080")
	log.Fatal(http.ListenAndServe(add, server.Router))
}
