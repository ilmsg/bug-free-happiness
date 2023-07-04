package main

import (
	"context"
	"fmt"

	"github.com/qiniu/qmgo"
)

var database *qmgo.Database
var collection *qmgo.Collection

func ConnectDatabase() {
	const databaseURI = "mongodb://localhost:27017"
	fmt.Println("Connecting to database:", databaseURI)

	ctx := context.Background()
	connection, err := qmgo.NewClient(ctx, &qmgo.Config{Uri: databaseURI})

	database = connection.Database("test")
	collection = database.Collection("books")

	defer func() {
		if err = connection.Close(ctx); err != nil {
			fmt.Println("Closing Connection to database", databaseURI)
			panic(err)
		}
	}()

}
