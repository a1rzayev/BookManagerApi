package routes

import (
	"book-manager-api/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.POST("/books", controllers.CreateBook)
	r.GET("/books", controllers.GetBooks)
	r.DELETE("/books/:id", controllers.DeleteBook)
}
