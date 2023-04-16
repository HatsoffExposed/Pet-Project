package routes

import (
	"petProject/pkg/controllers"

	"github.com/gin-gonic/gin"
)

var RegisterBookStoreRoutes = func(router *gin.Engine) {
	book := router.Group("/book")
	{
		book.POST("/", controllers.CreateBook)
		book.GET("/:bookId", controllers.GetBookById)
		book.GET("/", controllers.GetAllBooks)
		book.PUT("/:bookId", controllers.UpdateBook)
		book.DELETE("/:bookId", controllers.DeleteBook)
	}

}
