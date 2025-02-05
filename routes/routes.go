package routes

import (
	"book-manager-api/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	protected := r.Group("/")
	protected.Use(middlewares.AuthMiddleware())
	{
		protected.POST("/books", controllers.CreateBook)
		protected.GET("/books", controllers.GetBooks)
		protected.DELETE("/books/:id", controllers.DeleteBook)
	}
}
