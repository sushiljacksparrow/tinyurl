package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var DB *gorm.DB

func ConnectDataBase() {
	host := goDotEnvVariable("DB_HOST")
	dbPort := goDotEnvVariable("DB_PORT")
	dbName := goDotEnvVariable("DB_NAME")
	dbPassword := goDotEnvVariable("DB_PASSWORD")
	dbUser := goDotEnvVariable("DB_USER")
	database, err := gorm.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s",
		host, dbPort, dbUser, dbPassword, dbName))
	if err != nil {
		log.Panic("Failed to connect to DB", err)
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&URL{})

	DB = database
}


// use godot package to load/read the .env file and
// return the value of the key
func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}
