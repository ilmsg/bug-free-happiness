package main

import "github.com/gin-gonic/gin"

func getRouter() *gin.Engine {
	router := gin.Default()

	books := router.Group("/books")
	{
		books.GET("/", ListBook)
		books.POST("/", CreateBook)
		books.PUT("/", UpdateBook)
		books.DELETE("/", DeleteBook)
	}

	return router
}
