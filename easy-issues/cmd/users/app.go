package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/bug-free-happiness/easy-issues/application"
	"github.com/bug-free-happiness/easy-issues/domain"
	"github.com/bug-free-happiness/easy-issues/persistence/memory"
	"github.com/bug-free-happiness/easy-issues/web/controller"
)

func main() {
	userRepo := memory.NewUserRepository()
	userService := application.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	for i := 0; i < 10; i++ {
		userService.Create(&domain.User{Name: fmt.Sprintf("User_%d", i)})
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", userController.List)

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	server := &http.Server{
		Addr:           ":7003",
		Handler:        mux,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   10 * time.Second,
		IdleTimeout:    120 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(server.ListenAndServe())
}
