package main

import (
	"admin-panel/admin-panel/database"
	"admin-panel/admin-panel/initializers"
	"admin-panel/admin-panel/router"
	"log"
)

func init() {

	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("env not working ", err)

	}

	database.ConnectDB(&config)

}

func main() {

	r := router.SetupRouter()

	r.Run(":8080")
}
