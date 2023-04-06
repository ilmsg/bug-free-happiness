package main

import (
	"context"
	"net"

	pb "github.com/ilmsg/calculator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type CalServer struct {
	pb.CalServiceServer
}

func main() {
	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	pb.RegisterCalServiceServer(srv, &CalServer{})
	reflection.Register(srv)

	if err := srv.Serve(listener); err != nil {
		panic(err)
	}
}

func (s *CalServer) Add(ctx context.Context, request *pb.CalRequest) (*pb.CalResponse, error) {
	a, b := request.GetA(), request.GetB()
	result := a + b
	return &pb.CalResponse{Result: result}, nil
}

func (s *CalServer) Subtract(ctx context.Context, request *pb.CalRequest) (*pb.CalResponse, error) {
	a, b := request.GetA(), request.GetB()
	result := a - b
	return &pb.CalResponse{Result: result}, nil
}
