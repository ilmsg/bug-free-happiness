package main

import (
	"fmt"
	"os"

	"github.com/ilmsg/cafe/di/filedatabase"
	"github.com/ilmsg/cafe/di/runner"
)

func main() {
	db := filedatabase.NewFileDatabase("database.txt")
	runner := runner.NewRunner(db)
	if err := runner.Run(os.Stdout, os.Args); err != nil {
		fmt.Println(err)
	}
}
