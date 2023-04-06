package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	pb "github.com/ilmsg/calculator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:4040", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	client := pb.NewCalServiceClient(conn)

	app := gin.Default()

	app.GET("/add/:a/:b", func(ctx *gin.Context) {
		a, err := strconv.ParseUint(ctx.Param("a"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Parameter A"})
			return
		}

		b, err := strconv.ParseUint(ctx.Param("b"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Parameter B"})
			return
		}

		req := &pb.CalRequest{A: int64(a), B: int64(b)}
		if res, err := client.Add(ctx, req); err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"err": err.Error(),
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"result": fmt.Sprint(res.Result),
			})
		}
	})

	app.GET("/subtract/:a/:b", func(ctx *gin.Context) {
		a, err := strconv.ParseUint(ctx.Param("a"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Parameter A"})
			return
		}

		b, err := strconv.ParseUint(ctx.Param("b"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Parameter B"})
			return
		}

		req := &pb.CalRequest{A: int64(a), B: int64(b)}
		if res, err := client.Subtract(ctx, req); err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"err": err.Error(),
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"result": fmt.Sprint(res.Result),
			})
		}
	})

	if err := app.Run(":4050"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
