package main

import (
	"admin-panel/admin-panel/controllers"
	"admin-panel/admin-panel/database"
	"admin-panel/admin-panel/middlewares"
	"admin-panel/admin-panel/models"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	err := database.InitDatabase()
	if err != nil {

		log.Fatalln("could not create database", err)
	}

	database.GlobalDB.AutoMigrate(&models.User{})

	r := setupRouter()

	r.Run(":8080")
}

func setupRouter() *gin.Engine {

	r := gin.Default()

	r.GET("/eyupoglu/ilanlar", func(c *gin.Context) {
		c.String(200, "Welcome To This Website")
	})

	api := r.Group("/api")
	{

		public := api.Group("/public")
		{

			public.POST("/login", controllers.Login)

			// public.POST("/signup", controllers.Signup)
		}

		protected := api.Group("/protected").Use(middlewares.Authz())
		{

			protected.GET("/profile", controllers.Profile)
		}
	}

	return r
}
