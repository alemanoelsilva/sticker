package database

import (
	"log"
	stickers_entity "sticker/src/entities/stickers"

	users_entity "sticker/src/entities/users"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Instance *gorm.DB

var err error

func Connect(connectionString string) {
	Instance, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		panic("Cannot connect to DB")
	}
	log.Println("Connected to Database...")
}

func Migrate() {
	Instance.AutoMigrate(&stickers_entity.Sticker{}, &users_entity.User{})
	log.Println("Database Migration Completed...")
}
