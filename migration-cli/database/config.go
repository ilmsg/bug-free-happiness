package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func Open() *sql.DB {
	db, err := sql.Open("sqlite3", "mydb")
	if err != nil {
		log.Fatalf("Error Opening DB: %v\n", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error Pinging DB: %v\n", err)
	}

	fmt.Println("Connected to db!")
	return db
}
