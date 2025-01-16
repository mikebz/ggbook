package main

import (
	"os"
	"strconv"
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

func TestCreateGuestDx(t *testing.T) {
	p := map[string]any{
		"name":  "Test Guest",
		"email": "test@example.com",
	}
	rmap := createGuestDx(p)
	assert.NotNil(t, rmap)
	result := rmap["result"]
	assert.Equal(t, result, "OK")
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

func TestAllGuestsDx(t *testing.T) {
	createGuest(&Guest{Name: "Thing One", Email: "thing1@example.com"})
	createGuest(&Guest{Name: "Thing Two", Email: "thing2@example.com"})

	rmap := allGuestsDx(nil)
	assert.NotNil(t, rmap)
	result := rmap["result"]
	assert.Equal(t, result, "OK")
	assert.Greater(t, len(rmap), 5) // at least two guests with name and email and the result
}

func TestOneGuest(t *testing.T) {
	g := &Guest{Name: "Thing One", Email: "thing1@example.com"}
	createGuest(g)
	guest, err := oneGuest(g.ID)
	assert.NoError(t, err)
	assert.NotNil(t, guest)
}

func TestOneGuestDx(t *testing.T) {
	g := &Guest{Name: "Thing One", Email: "thing1@example.com"}
	createGuest(g)

	p := map[string]any{
		"id": strconv.FormatUint(uint64(g.ID), 10),
	}
	rmap := oneGuestDx(p)
	assert.NotNil(t, rmap)
	result := rmap["result"]
	assert.Equal(t, result, "OK")

}

func TestDeleteGuest(t *testing.T) {
	err := deleteGuest(1)
	assert.NoError(t, err)
}

func TestDeleteGuestDx(t *testing.T) {
	p := map[string]any{
		"id": "1",
	}
	rmap := deleteGuestDx(p)
	assert.NotNil(t, rmap)
	result := rmap["result"]
	assert.Equal(t, result, "OK")
}

func TestUpdateGuest(t *testing.T) {
	g := &Guest{Name: "Thing One", Email: "before@example.com"}
	err := createGuest(g)
	assert.NoError(t, err)

	g.Email = "after@example.com"
	err = updateGuest(g)
	assert.NoError(t, err)
}

func TestUpdateGuestDx(t *testing.T) {

	g := &Guest{Name: "Thing One", Email: "before@example.com"}
	err := createGuest(g)
	assert.NoError(t, err)

	p := map[string]any{
		"id": strconv.FormatUint(uint64(g.ID), 10),
		"name":  g.Name,
		"email": "after@example.com",
	}
	rmap := updateGuestDx(p)
	assert.NotNil(t, rmap)
	result := rmap["result"]
	assert.Equal(t, result, "OK")
}

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
