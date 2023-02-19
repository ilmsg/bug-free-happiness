package main

import (
	"log"

	"github.com/ilmsg/bug-free-happiness/gorm-seed/internal/models"
	"github.com/ilmsg/bug-free-happiness/gorm-seed/pkg/seeds"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	conn := OpenConnection()
	conn.AutoMigrate(&models.User{})
	for _, seed := range seeds.All() {
		if err := seed.Run(conn); err != nil {
			log.Fatalf("Running seed '%s', faiiled with error: %s", seed.Name, err.Error())
		}
	}
}

func OpenConnection() *gorm.DB {
	dsn := "orico:orico@tcp(127.0.0.1:33060)/orico?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}
