package main

import (
	"admin-panel/admin-panel/database"
	"admin-panel/admin-panel/initializers"
	"admin-panel/admin-panel/models"
	"fmt"
	"log"
)

func init() {

	config, err := initializers.LoadConfig("..")
	if err != nil {
		log.Fatal("env not working ", err)

	}

	database.ConnectDB(&config)

}
func main() {

	database.GlobalDB.AutoMigrate(&models.User{})
	database.GlobalDB.AutoMigrate(&models.Ilanlar{})
	fmt.Println("? migration complete")
}
