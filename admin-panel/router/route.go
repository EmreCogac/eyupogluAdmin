package router

import (
	"admin-panel/admin-panel/controllers"
	"admin-panel/admin-panel/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	r := gin.Default()

	r.GET("/eyupoglu/ilanlar", controllers.GetAll)
	// r.POST("/test", controllers.Delete)

	api := r.Group("/api")
	{

		public := api.Group("/public")
		{

			public.POST("/login", controllers.Login)

			// public.POST("/signup", controllers.Signup) // kanka kayıt işlemi için burayı kullan
		}

		protected := api.Group("/protected").Use(middlewares.Authz())
		{

			protected.GET("/profile", controllers.Profile)
			protected.POST("/create", controllers.CreatePost)
			protected.POST("/delete", controllers.Delete)
		}
	}

	return r
}
