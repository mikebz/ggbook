package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Create a new GORM model
type Guest struct {
	gorm.Model
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	db, err := gorm.Open(sqlite.Open("main.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Migration")
	err = db.AutoMigrate(&Guest{})
	if err != nil {
		panic("failed to migrate")
	}

	listenAndServe()
}
