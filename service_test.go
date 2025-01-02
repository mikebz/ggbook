package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setup() {
	openDb("test.db")
	migrate()
}

func TestCreateGuest(t *testing.T) {
	guest := &Guest{Name: "Test Guest", Email: "test@example.com"}
	err := createGuest(guest)
	if err != nil {
		t.Errorf("Error creating guest: %v", err)
	}
}

func TestAllGuests(t *testing.T) {
	createGuest(&Guest{Name: "Thing One", Email: "thing1@example.com"})
	createGuest(&Guest{Name: "Thing Two", Email: "thing2@example.com"})

	guests, err := allGuests()
	if err != nil {
		t.Errorf("Error getting all guests: %v", err)
	}

	assert.Greater(t, len(guests), 2)
}

func TestOneGuest(t *testing.T) {
	g := &Guest{Name: "Thing One", Email: "thing1@example.com"}
	createGuest(g)
	guest, err := oneGuest(g.ID)
	assert.NoError(t, err)
	assert.NotNil(t, guest)
}

func TestDeleteGuest(t *testing.T) {
	err := deleteGuest(1)
	assert.NoError(t, err)
}

// TODO: test updating a guest.

func teardown() {
	// not sure what happened to the gorm close method
	// but this is a test database, we should be OK
	os.Remove("test.db")
}

func TestMain(m *testing.M) {
	setup()
	defer teardown()
	code := m.Run()
	os.Exit(code)
}
