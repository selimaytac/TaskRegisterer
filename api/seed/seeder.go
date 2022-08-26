package seed

import (
	"github.com/jinzhu/gorm"
	"github.com/selimaytac/TaskRegisterer/api/models"
	"log"
	"time"
)

var users = []models.User{
	models.User{
		ID:                    0,
		Username:              "John Doe",
		Email:                 "john@gmail.com",
		Password:              "password",
		Department:            "Admin",
		CountOfTasks:          0,
		CountOfCompletedTasks: 0,
		IsActive:              false,
		IsDeleted:             false,
		CreatedAt:             time.Time{},
		UpdatedAt:             time.Time{},
	},
	models.User{
		ID:                    0,
		Username:              "Jane Doe",
		Email:                 "jame@gmail.com",
		Password:              "password",
		Department:            "Admin",
		CountOfTasks:          0,
		CountOfCompletedTasks: 0,
		IsActive:              false,
		IsDeleted:             false,
		CreatedAt:             time.Time{},
		UpdatedAt:             time.Time{},
	},
}

func Load(db *gorm.DB) {
	err := db.Debug().DropTableIfExists(&models.User{}).Error
	if err != nil {
		log.Fatalf("Cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.User{}).Error
	if err != nil {
		log.Fatalf("Cannot migrate table: %v", err)
	}

	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error

		if err != nil {
			log.Fatalf("Cannot seed users table: %v", err)
		}
	}
}
