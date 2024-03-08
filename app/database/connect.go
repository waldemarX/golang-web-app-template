package database

import (
	"fmt"
	"log"
	"strconv"

	"app/app/config"
	"app/app/internal"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Declare the variable for the database
var DB *gorm.DB

// ConnectDB connect to db
func ConnectDB() {
	var db_error error
	db_port := config.Config("DB_PORT")
	port, db_error := strconv.ParseUint(db_port, 10, 32)

	if db_error != nil {
		log.Println("Idiot")
	}

	// Connection URL to connect to Postgres Database
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Config("DB_HOST"),
		port,
		config.Config("DB_USER"),
		config.Config("DB_PASSWORD"),
		config.Config("DB_NAME"),
	)
	// Connect to the DB and initialize the DB variable
	DB, db_error = gorm.Open(postgres.Open(dsn))

	if db_error != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connection Opened to Database")

	// drop table at restart
	DB.Migrator().DropTable(&model.Note{})

	// Migrate the database
	DB.AutoMigrate(&model.Note{})
	fmt.Println("Database Migrated")
}
