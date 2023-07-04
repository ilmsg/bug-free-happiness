package main

import (
	"fmt"
)

func main() {
	router := getRouter()
	if err := router.Run(":8000"); err != nil {
		fmt.Println("Service is up & running at localhost:8000")
	}
}
