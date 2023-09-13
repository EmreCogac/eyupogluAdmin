package database

import (
	"admin-panel/admin-panel/initializers"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var GlobalDB *gorm.DB

func ConnectDB(config *initializers.Config) {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable ", config.DBHost, config.DBUserName, config.DBUserPassword, config.DBName, config.DBPort)
	GlobalDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the Database")
	}
	fmt.Println("? Connected Successfully to the Database")

}

func CloseDB() {

	db, err := GlobalDB.DB()
	if err != nil {
		log.Fatal("db couldnt close", err)
	}
	defer db.Close()
}
