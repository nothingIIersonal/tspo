package main

import (
	"pr7_1/services/auth"
	"pr7_1/services/mock_book"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/login", auth.Login)
	router.POST("/register", auth.Register)

	protected := router.Group("/")
	protected.Use(auth.AuthMiddleware())
	{
		protected.GET("/books", mock_book.GetBooks)
		protected.POST("/books", mock_book.CreateBook)
		protected.PUT("/books", mock_book.UpdateBook)
		protected.DELETE("/books", mock_book.DeleteBook)
	}

	router.Run(":8080")
}
