package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/ilmsg/ruslan-tsyganok/grpc/gen/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:3001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println(err)
	}

	client := pb.NewTestApiClient(conn)
	resp, err := client.Echo(context.Background(), &pb.ResponseRequest{Msg: "Hello World"})
	if err != nil {
		log.Println(err)
	}

	fmt.Println(resp)
	fmt.Println(resp.Msg)

	respUser, err := client.GetUser(context.Background(), &pb.UserRequest{Uuid: "001"})
	if err != nil {
		log.Println(err)
	}

	fmt.Println(respUser)
	fmt.Println(respUser.Email, respUser.Name, respUser.Age)
}
