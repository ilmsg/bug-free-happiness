package main

import "github.com/gin-gonic/gin"

type Book struct {
}

type BookResponse struct {
}

type BookCreateUpdateReqeust struct {
}

func ListBook(ctx *gin.Context) {
	ctx.JSON(2000, gin.H{
		"message": "created",
	})
}

func CreateBook(ctx *gin.Context) {
	ctx.JSON(2000, gin.H{
		"message": "created",
	})
}

func UpdateBook(ctx *gin.Context) {
	ctx.JSON(2000, gin.H{
		"message": "created",
	})
}

func DeleteBook(ctx *gin.Context) {
	ctx.JSON(2000, gin.H{
		"message": "created",
	})
}

func GetBooksResponse(book Book) BookResponse {
	bookResponse := BookResponse{}
	return bookResponse
}
