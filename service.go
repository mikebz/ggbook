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

var db *gorm.DB

// the database string is empty we will default to `main.db`
func openDb(dbUrl string) error {
	var err error

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

func createGuest(guest *Guest) error {
	fmt.Println("Creating user")
	return db.Create(guest).Error
}

func allGuests() (guests []Guest, err error) {
	fmt.Println("Getting all users")
	err = db.Find(&guests).Error
	return guests, err
}

func oneGuest(id uint) (guest *Guest, err error) {
	fmt.Println("Getting one user")
	err = db.First(&guest, id).Error
	return guest, err
}

func deleteGuest(id uint) error {
	fmt.Println("Deleting user")
	err := db.Delete(&Guest{}, id).Error
	return err
}

func updateGuest(guest *Guest) error {
	fmt.Println("Updating user")
	err := db.Model(guest).Updates(guest).Error
	return err
}
