package db

import (
	"github.com/arunnasarain/todo/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(url string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.TODO{})

	return db
}

func InitialData(db *gorm.DB) {
	todos := []models.TODO{
		{Title: "Task 1", Content: "TODO1"},
		{Title: "Task 2", Content: "TODO1"},
		{Title: "Task 3", Content: "TODO1"},
	}

	for _, todo := range todos {
		result := db.Create(&todo)
		if result.Error != nil {
			log.Fatal("Failed to add todo:", result.Error)
		}
	}
}
