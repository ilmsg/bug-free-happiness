package main

import (
	"context"
	"log"
	"net"

	pb "github.com/ilmsg/ruslan-tsyganok/grpc/gen/proto"
	"google.golang.org/grpc"
)

type testApiServer struct {
	pb.UnimplementedTestApiServer
}

func (s *testApiServer) GetUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	return &pb.UserResponse{Name: "Scott Tiger", Age: 38, Email: "scott.tiger@gmail.com"}, nil
}

func (s *testApiServer) Echo(ctx context.Context, req *pb.ResponseRequest) (*pb.ResponseRequest, error) {
	// return &pb.ResponseRequest{}, nil
	return req, nil
}

func main() {
	listen, err := net.Listen("tcp", "localhost:3001")
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterTestApiServer(grpcServer, &testApiServer{})

	grpcServer.Serve(listen)
}
