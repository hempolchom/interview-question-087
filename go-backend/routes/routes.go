package routes

import (
	"example.com/interview-question-087/go-backend/controllers"
	"example.com/interview-question-087/go-backend/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.CORSMiddleware())

	api := r.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/register", controllers.Register)
			auth.POST("/login", controllers.Login)

			protected := auth.Group("")
			protected.Use(middleware.AuthMiddleware())
			{
				protected.GET("/profile", controllers.GetProfile)
			}
		}
	}

	return r
}