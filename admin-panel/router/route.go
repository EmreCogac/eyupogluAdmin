package router

import (
	"admin-panel/admin-panel/controllers"
	"admin-panel/admin-panel/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	r := gin.Default()

	r.GET("/eyupoglu/ilanlar", controllers.GetAll)

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
