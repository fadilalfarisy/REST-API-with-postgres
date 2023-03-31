package routers

import (
	"challenge-2/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.GET("/books", controllers.GetAllBooks)
	router.GET("/book/:bookID", controllers.GetBookById)
	router.POST("/book", controllers.CreateBook)
	router.PUT("/book/:bookID", controllers.UpdateBooks)
	router.DELETE("/book/:bookID", controllers.DeleteBook)

	return router
}
