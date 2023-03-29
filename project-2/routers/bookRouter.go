package routers

import (
	"project-2/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.GET("/books", controllers.GetBooks)
	router.GET("/books/:id", controllers.GetBook)
	router.POST("/books", controllers.AddBook)
	router.PUT("/books/:id", controllers.UpdateBook)
	router.DELETE("/books/:id", controllers.DeleteBook)

	return router
}
