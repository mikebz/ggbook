package main

import (
	"fmt"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Create a new GORM model
type Guest struct {
	gorm.Model
	Name  string `json:"name"`
	Email string `json:"email"`
}

var db *gorm.DB

func openDb() error {
	var err error
	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		dbUrl = "main.db"
	}

	fmt.Println("Opening database")
	db, err = gorm.Open(sqlite.Open(dbUrl), &gorm.Config{})
	return err
}

func migrate() error {
	fmt.Println("Migration")
	err := db.AutoMigrate(&Guest{})
	return err
}
