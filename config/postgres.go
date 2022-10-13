package config

import (
	"edufunds/models"
	"fmt"
	"os"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/joho/godotenv"
)
const Dbdriver ="postgres"
func SetUpDatabaseConnection() *gorm.DB {
	errEnv := godotenv.Load()

	if errEnv != nil {
		panic("Failed to load env file")
	}
	// dbPORT := os.Getenv("APP_PORT")
	dbUser := os.Getenv("POSTGRES_USER")
	dbPass := os.Getenv("POSTGRES_PASSWORD")
	dbHost := os.Getenv("POSTGRES_ADDRESS")
	dbName := os.Getenv("POSTGRES_DB")
	dbPort := os.Getenv("POSTGRES_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",dbHost,dbUser,dbPass,dbName,dbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		// panic("Failed to create a connection to the DB")
		fmt.Println("ini errorrrnya >>>>>>>",err)
	}
	
	//Migrate
	db.AutoMigrate(&models.User{})
	return db
}


func CloseDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()

	if err != nil {
		panic("Failed to close connection from the database")
	}

	dbSQL.Close()
}
