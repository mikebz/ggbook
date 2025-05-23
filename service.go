package main

import (
	"fmt"
	"strconv"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Create a new GORM model
type Guest struct {
	gorm.Model
	Name  string `json:"name"`
	Email string `json:"email"`
}

type DxFunc func(map[string]any) map[string]any

var db *gorm.DB

// the database string is empty we will default to `main.db`
func openDb(dbUrl string) error {
	var err error

	if dbUrl == "" {
		dbUrl = "main.db"
	}

	logger.Println("Opening database")
	db, err = gorm.Open(sqlite.Open(dbUrl), &gorm.Config{})
	return err

}

func migrate() error {
	logger.Println("Migration")
	err := db.AutoMigrate(&Guest{})
	return err
}

func createGuestDx(p map[string]any) map[string]any {
	guest := &Guest{Name: p["name"].(string), Email: p["email"].(string)}
	err := createGuest(guest)
	if err != nil {
		return map[string]any{
			"result": "error: " + err.Error(),
		}
	} else {
		return map[string]any{
			"result": "OK",
		}
	}
}

func createGuest(guest *Guest) error {
	logger.Println("Creating user")
	return db.Create(guest).Error
}

func allGuestsDx(map[string]any) map[string]any {
	allguests, err := allGuests()
	if err != nil {
		return map[string]any{
			"result": "error: " + err.Error(),
		}
	}

	result := make(map[string]any)

	// here we are trying to flatten structs into a map
	for g := range allguests {
		idStr := strconv.FormatUint(uint64(allguests[g].ID), 10)
		result["guest["+idStr+"].name"] = allguests[g].Name
		result["guest["+idStr+"].email"] = allguests[g].Email
		result["guest["+idStr+"].id"] = allguests[g].ID
	}

	result["result"] = "OK"
	return result
}

func allGuests() (guests []Guest, err error) {
	logger.Println("Getting all users")
	err = db.Find(&guests).Error
	logger.Printf("Got %d guests\n", len(guests))
	return guests, err
}

func oneGuestDx(p map[string]any) map[string]any {
	logger.Printf("Getting one user, params %v", p)

	id, err := safeUint(p["id"])
	if err != nil {
		return map[string]any{
			"result": "error: " + err.Error(),
		}
	}

	g, err := oneGuest(uint(id))
	if err != nil {
		return map[string]any{
			"result": "error: " + err.Error(),
		}
	}

	gmap := guestMap(g)
	gmap["result"] = "OK"
	return gmap
}

func guestMap(guest *Guest) map[string]any {
	return map[string]any{
		"id":    guest.ID,
		"name":  guest.Name,
		"email": guest.Email,
	}
}

func oneGuest(id uint) (guest *Guest, err error) {
	logger.Println("Getting one user")
	err = db.First(&guest, id).Error
	return guest, err
}

func deleteGuestDx(p map[string]any) map[string]any {
	logger.Printf("Deleting user, params %v", p)
	id, err := safeUint(p["id"])
	if err != nil {
		return map[string]any{
			"result": "error: " + err.Error(),
		}
	}

	err = deleteGuest(uint(id))

	if err != nil {
		return map[string]any{
			"result": "error: " + err.Error(),
		}
	} else {
		return map[string]any{
			"result": "OK",
		}
	}
}

func deleteGuest(id uint) error {
	logger.Println("Deleting user")
	err := db.Delete(&Guest{}, id).Error
	return err
}

func updateGuestDx(p map[string]any) map[string]any {
	logger.Printf("Updating user, params %v", p)
	id, err := safeUint(p["id"])
	if err != nil {
		return map[string]any{
			"result": "error: " + err.Error(),
		}
	}

	guest := &Guest{Name: p["name"].(string), Email: p["email"].(string)}
	guest.ID = uint(id)
	err = updateGuest(guest)

	if err != nil {
		return map[string]any{
			"result": "error: " + err.Error(),
		}
	} else {
		return map[string]any{
			"result": "OK",
		}
	}
}

func updateGuest(guest *Guest) error {
	logger.Println("Updating user")
	err := db.Model(guest).Updates(guest).Error
	return err
}

func safeUint(value interface{}) (uint, error) {
	if value == nil {
		return 0, fmt.Errorf("value is nil")
	}

	switch v := value.(type) {
	case uint:
		return v, nil
	case int:
		if v < 0 {
			return 0, fmt.Errorf("cannot convert negative int to uint")
		}
		return uint(v), nil
	case string:
		u, err := strconv.ParseUint(v, 10, 64)
		if err != nil {
			return 0, fmt.Errorf("invalid uint string: %w", err)
		}
		return uint(u), nil

	case float32, float64:
		f := value.(float64)
		if f < 0 {
			return 0, fmt.Errorf("cannot convert negative float to uint")
		}
		return uint(f), nil

	default:
		return 0, fmt.Errorf("unsupported type: %T", value)
	}
}