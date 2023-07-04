package main

import (
	"fmt"
	"time"
)

type Server struct {
	quitch chan struct{}
}

func (s *Server) start() {
free:
	for {
		select {
		case <-s.quitch:
			fmt.Println("attempting graceful shutdown")
			break free
		default:
		}
	}
}

func (s *Server) quit() {
	close(s.quitch)
}

func main() {
	s := &Server{
		quitch: make(chan struct{}),
	}

	go func() {
		time.Sleep(1 * time.Second)
		s.quit()
	}()

	s.start()
}
