package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var DB *gorm.DB

func ConnectDataBase() {
	database, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		log.Panic("Failed to connect to DB", err)
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&URL{})

	DB = database
}
