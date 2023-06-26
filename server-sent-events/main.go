package main

import (
	"embed"
	"html/template"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

//go:embed assets/* templates/*
var f embed.FS
var max = 3
var min = 1

func main() {
	router := gin.Default()
	templ := template.Must(template.New("").ParseFS(f, "templates/*.tpl"))
	router.SetHTMLTemplate(templ)
	router.StaticFS("/public", http.FS(f))

	// SSE endpoint
	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.tpl", nil)
	})
	router.GET("/progress", progressor)

	// Start the server
	if err := router.Run(":3000"); err != nil {
		panic(err)
	}
}

func progressor(ctx *gin.Context) {
	noOfExecution := 10
	// noOfExecution := rand.Intn(max-min) + min
	progress := 0
	for progress <= noOfExecution {
		progressPercentage := float64(progress) / float64(noOfExecution) * 100

		ctx.SSEvent("progress", map[string]interface{}{
			"currentTask":        progress,
			"progressPercentage": progressPercentage,
			"noOftasks":          noOfExecution,
			"completed":          false,
		})
		// Flush the response to ensure the data is sent immediately
		ctx.Writer.Flush()

		// progress += 1
		progress += rand.Intn(max-min) + min
		time.Sleep(1 * time.Second)
	}

	ctx.SSEvent("progress", map[string]interface{}{
		"completed":          true,
		"progressPercentage": 100,
	})

	// Flush the response to ensure the data is sent immediately
	ctx.Writer.Flush()

}
